package action

import (
	"context"
	"fmt"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/registry/storage"
)

const (
	actionRegistryName = "/actions/"
)

// Storage defines the storage for Action objects
type Storage struct {
	storage storage.Storage
	destroy storage.DestroyFunc
	codec   codec.Codec
}

// NewActionStorage creates a new storage for Action objects
func NewActionStorage(storage storage.Storage, codec codec.Codec, destroy storage.DestroyFunc) *Storage {
	return &Storage{storage: storage, codec: codec, destroy: destroy}
}

// Close closes the storage connection
func (s *Storage) Close() error {
	s.destroy()
	return s.storage.Close()
}

// Create creates a new Action
func (s *Storage) Create(ctx context.Context, a *models.Action) error {
	data, err := s.codec.Encode(a)
	if err != nil {
		return err
	}
	return s.storage.Create(
		ctx,
		fmt.Sprintf("%s%s/%s", actionRegistryName, a.Module, a.Name),
		data,
		0)
}

// Update updates an Action
func (s *Storage) Update(ctx context.Context, a *models.Action) error {
	data, err := s.codec.Encode(a)
	if err != nil {
		return err
	}
	return s.storage.Update(
		ctx,
		fmt.Sprintf("%s%s/%s", actionRegistryName, a.Module, a.Name),
		data,
	)
}

// Delete deletes an Action object
func (s *Storage) Delete(ctx context.Context, module, name string) error {
	return s.storage.Delete(
		ctx,
		fmt.Sprintf("%s%s/%s", actionRegistryName, module, name),
	)
}

// Get gets an action
func (s *Storage) Get(ctx context.Context, module, name string) (*models.Action, error) {
	data, err := s.storage.Get(ctx, fmt.Sprintf("%s%s/%s", actionRegistryName, module, name))
	if err != nil {
		return nil, err
	}

	var a models.Action
	err = s.codec.Decode(data, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// List returns all actions
func (s *Storage) List(ctx context.Context) ([]*models.Action, error) {
	data, err := s.storage.List(ctx, actionRegistryName)
	if err != nil {
		return nil, err
	}
	var al []*models.Action
	for _, d := range data {
		var a models.Action
		err = s.codec.Decode(d, &a)
		if err != nil {
			return nil, err
		}
		al = append(al, &a)
	}

	return al, nil
}

// Watch watches for actions
func (s *Storage) Watch(ctx context.Context) (storage.Watcher, error) {
	return s.storage.Watch(ctx, actionRegistryName)
}
