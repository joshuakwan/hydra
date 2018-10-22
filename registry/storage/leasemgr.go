package storage

import (
	"context"
	"log"
	"sync"

	"go.etcd.io/etcd/clientv3"
)

//var leaseMgr *leaseManager

type leaseManager struct {
	client  *clientv3.Client
	leaseMu sync.Mutex
	// currentLeaseID clientv3.LeaseID
	// currentTTL     int64
}

func newLeaseManager(client *clientv3.Client) *leaseManager {
	return &leaseManager{client: client}
}

func (l *leaseManager) GetLease(ctx context.Context, ttl int64) clientv3.LeaseID {
	l.leaseMu.Lock()
	defer l.leaseMu.Unlock()

	// if ttl == l.currentTTL {
	// 	return l.currentLeaseID
	// }

	lresp, err := l.client.Lease.Grant(ctx, ttl)

	if err != nil {
		log.Println(err)
		return clientv3.LeaseID(0)
	}
	// l.currentTTL = ttl
	// l.currentLeaseID = clientv3.LeaseID(lresp.ID)
	return clientv3.LeaseID(lresp.ID)
}
