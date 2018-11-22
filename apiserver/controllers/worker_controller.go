package controllers

import (
	"github.com/joshuakwan/hydra/apiserver/services"
	"github.com/joshuakwan/hydra/models"
	"github.com/kataras/iris"
)

// WorkerController handles HTTP requests on /workers
type WorkerController struct {
	Service services.WorkerService
}

// Get handles GET /workers
func (w *WorkerController) Get() (results []*models.Worker) {
	return w.Service.List()
}

// GetBy handles GET /workers/<name>
func (w *WorkerController) GetBy(name string) (worker *models.Worker, found bool) {
	return w.Service.Get(name)
}

// Post handles POST /workers
func (w *WorkerController) Post(ctx iris.Context) error {
	var worker models.Worker
	err := ctx.ReadJSON(&worker)
	if err != nil {
		return err
	}
	return w.Service.Register(&worker)
}

// PostReport handles POST /workers/report
func (w *WorkerController) PostReport(ctx iris.Context) error {
	var worker models.Worker
	err := ctx.ReadJSON(&worker)
	if err != nil {
		return err
	}
	return w.Service.Report(&worker)
}
