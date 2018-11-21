package services

import (
	"context"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/registry/worker"
)

// WorkerService defines the REST interface for Workers
type WorkerService interface {
	Register(worker *models.Worker) error
	List() []*models.Worker
	Get(name string) (*models.Worker, bool)
}

// NewWorkerService initializes a concrete ActionService implementation
func NewWorkerService() WorkerService {
	storage, err := worker.NewWorkerStorage(codec.NewCodec("json"))
	if err != nil {
		return nil
	}
	return &workerService{storage: storage}
}

type workerService struct {
	storage *worker.Storage
}

func (w *workerService) Register(worker *models.Worker) error {
	return w.storage.Create(context.Background(), worker)
}

func (w *workerService) List() []*models.Worker {
	wl, err := w.storage.List(context.Background())
	if err != nil {
		return nil
	}
	return wl
}

func (w *workerService) Get(name string) (*models.Worker, bool) {
	worker, err := w.storage.Get(context.Background(), name)
	if err != nil {
		return nil, false
	}
	return worker, true
}
