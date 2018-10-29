package rule

import (
	"context"
	"fmt"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	store "github.com/joshuakwan/hydra/registry/storage"
)

const (
	ruleRegistryName = "/rules/"
)

// Storage defines the storage for Rule objects
type Storage struct {
	storage store.Storage
	destroy store.DestroyFunc
	codec   codec.Codec
}

// NewRuleStorage creates a new storage for Rule objects
func NewRuleStorage(codec codec.Codec) (*Storage, error) {
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

// Create creates a new Rule
func (s *Storage) Create(ctx context.Context, r *models.Rule) error {
	data, err := s.codec.Encode(r)
	if err != nil {
		return err
	}
	return s.storage.Create(
		ctx,
		fmt.Sprintf("%s%s/%s", ruleRegistryName, r.Module, r.Name),
		data,
		0)
}

// Update updates an Action
func (s *Storage) Update(ctx context.Context, r *models.Action) error {
	data, err := s.codec.Encode(r)
	if err != nil {
		return err
	}
	return s.storage.Update(
		ctx,
		fmt.Sprintf("%s%s/%s", ruleRegistryName, r.Module, r.Name),
		data,
	)
}

// Delete deletes an Action object
func (s *Storage) Delete(ctx context.Context, module, name string) error {
	return s.storage.Delete(
		ctx,
		fmt.Sprintf("%s%s/%s", ruleRegistryName, module, name),
	)
}

// Get gets an action
func (s *Storage) Get(ctx context.Context, module, name string) (*models.Rule, error) {
	data, err := s.storage.Get(ctx, fmt.Sprintf("%s%s/%s", ruleRegistryName, module, name))
	if err != nil {
		return nil, err
	}

	var r models.Rule
	err = s.codec.Decode(data, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// List returns all actions
func (s *Storage) List(ctx context.Context) ([]*models.Rule, error) {
	data, err := s.storage.List(ctx, ruleRegistryName)
	if err != nil {
		return nil, err
	}
	var rl []*models.Rule
	for _, d := range data {
		var r models.Rule
		err = s.codec.Decode(d, &r)
		if err != nil {
			return nil, err
		}
		rl = append(rl, &r)
	}

	return rl, nil
}

// Watch watches for actions
func (s *Storage) Watch(ctx context.Context) (store.Watcher, error) {
	return s.storage.Watch(ctx, ruleRegistryName, s.codec)
}
