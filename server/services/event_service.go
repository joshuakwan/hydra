package services

import (
	"context"

	"github.com/joshuakwan/hydra/registry/storage"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/registry/event"
)

type EventService interface {
	Watch() storage.Watcher
}

func NewEventService() EventService {
	storage, err := event.NewEventStorage(codec.NewCodec("json"))
	if err != nil {
		return nil
	}
	return &eventService{storage: storage}
}

type eventService struct {
	storage *event.Storage
}

func (e *eventService) Watch() storage.Watcher {
	watch, err := e.storage.Watch(context.Background())
	if err != nil {
		return nil
	}
	return watch
}
