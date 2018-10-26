package services

import (
	"context"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/registry/action"
)

type ActionService interface {
	Create(action *models.Action) error
	List() []*models.Action
	Get(module, name string) (*models.Action, bool)
}

func NewActionService() ActionService {
	storage, err := action.NewActionStorage(codec.NewCodec("json"))
	if err != nil {
		return nil
	}
	return &actionService{storage: storage}
}

type actionService struct {
	storage *action.Storage
}

func (a *actionService) Create(action *models.Action) error {
	return a.storage.Create(context.Background(), action)
}

func (a *actionService) List() []*models.Action {
	al, err := a.storage.List(context.Background())
	if err != nil {
		return nil
	}
	return al
}

func (a *actionService) Get(module, name string) (*models.Action, bool) {
	action, err := a.storage.Get(context.Background(), module, name)
	if err != nil {
		return nil, false
	}
	return action, true
}
