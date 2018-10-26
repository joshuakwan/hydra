package controllers

import (
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/server/services"
	"github.com/kataras/iris"
)

// ActionController handles HTTP requests on /actions
type ActionController struct {
	Service services.ActionService
}

// Get handles GET /actions
func (a *ActionController) Get() (results []*models.Action) {
	return a.Service.List()
}

// GetBy handles GET /actions/<module>/<name>
func (a *ActionController) GetBy(module, name string) (action *models.Action, found bool) {
	return a.Service.Get(module, name)
}

// Post handles POST /actions
func (a *ActionController) Post(ctx iris.Context) error {
	var action models.Action
	err := ctx.ReadJSON(&action)
	if err != nil {
		return err
	}
	return a.Service.Create(&action)
}
