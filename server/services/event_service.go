package services

import (
	"context"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/registry/event"
	"github.com/joshuakwan/hydra/registry/storage"
)

// EventService defines the REST interface for Events
type EventService interface {
	List() []*models.Event
	Create(event *models.Event) error
	Watch() storage.Watcher
}

// NewEventService initializes a concrete EventService implementation
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

func (e *eventService) Create(event *models.Event) error {
	return e.storage.Create(context.Background(), event)
}

func (e *eventService) List() []*models.Event {
	el, err := e.storage.List(context.Background())
	if err != nil {
		return nil
	}
	return el
}

func (e *eventService) Watch() storage.Watcher {
	watch, err := e.storage.Watch(context.Background())
	if err != nil {
		return nil
	}
	return watch
}
