package main

import (
	"github.com/joshuakwan/hydra/server/controllers"
	"github.com/joshuakwan/hydra/server/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func actions(app *mvc.Application) {
	actionService := services.NewActionService()
	app.Register(actionService)
	app.Handle(new(controllers.ActionController))
}

func newApp() *iris.Application {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	mvc.Configure(app.Party("/actions"), actions)

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
