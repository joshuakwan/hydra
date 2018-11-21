package worker

import (
	"context"
	"fmt"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	store "github.com/joshuakwan/hydra/registry/storage"
)

const (
	workerRegistryName = "/workers/"
)

// Storage defines the storage for Action objects
type Storage struct {
	storage store.Storage
	destroy store.DestroyFunc
	codec   codec.Codec
}

// NewWorkerStorage creates a new storage for Worker objects
func NewWorkerStorage(codec codec.Codec) (*Storage, error) {
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

// Create creates a new Worker
func (s *Storage) Create(ctx context.Context, w *models.Worker) error {
	data, err := s.codec.Encode(w)
	if err != nil {
		return err
	}
	return s.storage.Create(
		ctx,
		fmt.Sprintf("%s%s", workerRegistryName, w.Name),
		data,
		0)
}

// Update updates a Worker
func (s *Storage) Update(ctx context.Context, w *models.Worker) error {
	data, err := s.codec.Encode(w)
	if err != nil {
		return err
	}
	return s.storage.Update(
		ctx,
		fmt.Sprintf("%s%s", workerRegistryName, w.Name),
		data,
	)
}

// Delete deletes a Worker object
func (s *Storage) Delete(ctx context.Context, name string) error {
	return s.storage.Delete(
		ctx,
		fmt.Sprintf("%s%s", workerRegistryName, name),
	)
}

// Get gets an Worker
func (s *Storage) Get(ctx context.Context, name string) (*models.Worker, error) {
	data, err := s.storage.Get(ctx, fmt.Sprintf("%s%s", workerRegistryName, name))
	if err != nil {
		return nil, err
	}

	var w models.Worker
	err = s.codec.Decode(data, &w)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

// List returns all Workers
func (s *Storage) List(ctx context.Context) ([]*models.Worker, error) {
	data, err := s.storage.List(ctx, workerRegistryName)
	if err != nil {
		return nil, err
	}
	var wl []*models.Worker
	for _, d := range data {
		var w models.Worker
		err = s.codec.Decode(d, &w)
		if err != nil {
			return nil, err
		}
		wl = append(wl, &w)
	}

	return wl, nil
}

// Watch watches for Workers
func (s *Storage) Watch(ctx context.Context) (store.Watcher, error) {
	return s.storage.Watch(ctx, workerRegistryName, s.codec)
}
