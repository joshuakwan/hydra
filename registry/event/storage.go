package event

import (
	"context"
	"fmt"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	store "github.com/joshuakwan/hydra/registry/storage"
)

const (
	eventRegistryName = "/events/"
)

// Storage defines the storage for Event objects
type Storage struct {
	storage store.Storage
	destroy store.DestroyFunc
	codec   codec.Codec
}

// NewEventStorage creates a new storage for Event objects
func NewEventStorage(codec codec.Codec) (*Storage, error) {
	storage, destroy, err := store.NewStorage()
	if err != nil {
		return nil, err
	}
	return &Storage{storage: storage, codec: codec, destroy: destroy}, nil
}

// Close closes the storage connection
func (s *Storage) Close() error {
	s.destroy()
	return s.storage.Close()
}

// Create creates a new Event
func (s *Storage) Create(ctx context.Context, e *models.Event) error {
	data, err := s.codec.Encode(e)
	if err != nil {
		return err
	}
	return s.storage.Create(
		ctx,
		fmt.Sprintf("%s%s/%d", eventRegistryName, e.Source, e.Timestamp),
		data,
		60)
}

// Delete deletes an Event object
func (s *Storage) Delete(ctx context.Context, source string, timestamp int64) error {
	return s.storage.Delete(
		ctx,
		fmt.Sprintf("%s%s/%d", eventRegistryName, source, timestamp),
	)
}

// Get gets an Event
func (s *Storage) Get(ctx context.Context, source string, timestamp int64) (*models.Event, error) {
	data, err := s.storage.Get(ctx, fmt.Sprintf("%s%s/%d", eventRegistryName, source, timestamp))
	if err != nil {
		return nil, err
	}

	var e models.Event
	err = s.codec.Decode(data, &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// List returns all Events
func (s *Storage) List(ctx context.Context) ([]*models.Event, error) {
	data, err := s.storage.List(ctx, eventRegistryName)
	if err != nil {
		return nil, err
	}
	var el []*models.Event
	for _, d := range data {
		var e models.Event
		err = s.codec.Decode(d, &e)
		if err != nil {
			return nil, err
		}
		el = append(el, &e)
	}

	return el, nil
}

// Watch watches for Events
func (s *Storage) Watch(ctx context.Context) (store.Watcher, error) {
	return s.storage.Watch(ctx, eventRegistryName, s.codec)
}
