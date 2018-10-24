package storage

import (
	"context"
	"log"
	"sync"

	"go.etcd.io/etcd/clientv3"
)

type watcher struct {
	client *clientv3.Client
}

type watchChan struct {
	watcher           *watcher
	key               string
	ctx               context.Context
	cancel            context.CancelFunc
	incomingEventChan chan *clientv3.Event
	resultChan        chan WatchEvent
	errChan           chan error
}

func (w *watcher) Watch(ctx context.Context, key string) (Watcher, error) {
	wc := w.buildWatchChan(ctx, key)
	go wc.run()
	return wc, nil
}

// TODO extract the buffer size to configuration
func (w *watcher) buildWatchChan(ctx context.Context, key string) *watchChan {
	wc := &watchChan{
		watcher:           w,
		key:               key,
		incomingEventChan: make(chan *clientv3.Event, 10),
		resultChan:        make(chan WatchEvent, 10),
		errChan:           make(chan error, 1),
	}
	wc.ctx, wc.cancel = context.WithCancel(ctx)
	return wc
}

func (wc *watchChan) ResultChan() <-chan WatchEvent {
	return wc.resultChan
}

func (wc *watchChan) Stop() {
	wc.cancel()
}

func (wc *watchChan) run() {
	watchClosedCh := make(chan struct{})

	// inbound logic
	go wc.startWatching(watchClosedCh)

	// outbound logic
	var resultChanWG sync.WaitGroup
	resultChanWG.Add(1)
	go wc.processEvent(&resultChanWG)

	select {
	case err := <-wc.errChan:
		if err == context.Canceled {
			break
		}
		if err != nil {
			select {
			case wc.resultChan <- wc.errorToWatchEvent(err):
			case <-wc.ctx.Done():
			}
		}
	case <-watchClosedCh:
	case <-wc.ctx.Done():
	}

	wc.cancel()
	resultChanWG.Wait()
	close(wc.resultChan)
}

func (wc *watchChan) startWatching(watchClosedCh chan struct{}) {
	log.Printf("key = %s", wc.key)
	wch := wc.watcher.client.Watch(wc.ctx, wc.key, clientv3.WithPrefix())
	for wresp := range wch {
		if wresp.Err() != nil {
			wc.sendError(wresp.Err())
			return
		}
		for _, e := range wresp.Events {
			wc.sendEvent(e)
		}
	}
	close(watchClosedCh)
}

func (wc *watchChan) processEvent(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case e := <-wc.incomingEventChan:
			select {
			case wc.resultChan <- wc.eventToWatchEvent(e):
			case <-wc.ctx.Done():
				return
			}
		case <-wc.ctx.Done():
			return
		}
	}
}

func (wc *watchChan) sendError(err error) {
	select {
	case wc.errChan <- err:
	case <-wc.ctx.Done():
	}
}

func (wc *watchChan) sendEvent(e *clientv3.Event) {
	select {
	case wc.incomingEventChan <- e:
	case <-wc.ctx.Done():
	}
}

func (wc *watchChan) errorToWatchEvent(err error) WatchEvent {
	return WatchEvent{Type: Error, Data: nil}
}

func (wc *watchChan) eventToWatchEvent(e *clientv3.Event) WatchEvent {
	var we WatchEvent

	if e.IsCreate() {
		we.Type = Created
	} else if e.IsModify() {
		we.Type = Updated
	} else if e.Type.String() == "DELETE" {
		we.Type = Deleted
	} else {
		return WatchEvent{Type: Error, Data: nil}
	}

	we.Data = e.Kv.Value

	return we
}
