package controllers

import (
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/server/services"
	"github.com/kataras/iris"
)

// RuleController handles HTTP requests on /rules
type RuleController struct {
	Service services.RuleService
}

// Get handles GET /rules
func (r *RuleController) Get() (results []*models.Rule) {
	return r.Service.List()
}

// GetBy handles GET /rules/<module>/<name>
func (r *RuleController) GetBy(module, name string) (rule *models.Rule, found bool) {
	return r.Service.Get(module, name)
}

// Post handles POST /rules
func (r *RuleController) Post(ctx iris.Context) error {
	var rule models.Rule
	err := ctx.ReadJSON(&rule)
	if err != nil {
		return err
	}
	return r.Service.Create(&rule)
}
