package main

import (
	"github.com/joshuakwan/hydra/apiserver/controllers"
	"github.com/joshuakwan/hydra/apiserver/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func actions(app *mvc.Application) {
	actionService := services.NewActionService()
	app.Register(actionService)
	app.Handle(new(controllers.ActionController))
}

func events(app *mvc.Application) {
	eventService := services.NewEventService()
	app.Register(eventService)
	app.Handle(new(controllers.EventController))
}

func rules(app *mvc.Application) {
	ruleService := services.NewRuleService()
	app.Register(ruleService)
	app.Handle(new(controllers.RuleController))
}

func workers(app *mvc.Application) {
	workerService := services.NewWorkerService()
	app.Register(workerService)
	app.Handle(new(controllers.WorkerController))
}

func newApp() *iris.Application {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	mvc.Configure(app.Party("/actions"), actions)
	mvc.Configure(app.Party("/events"), events)
	mvc.Configure(app.Party("/rules"), rules)
	mvc.Configure(app.Party("/workers"), workers)

	return app
}

func main() {
	app := newApp()
	app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
