package controllers

import (
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/server/services"
	"github.com/kataras/iris"
)

type EventController struct {
	Service services.EventService
}

func (e *EventController) GetWatch(ctx iris.Context) {
	ctx.ContentType("application/json")
	ctx.Header("Transfer-Encoding", "chunked")

	watch := e.Service.Watch()
	defer watch.Stop()
	for {
		for w := range watch.ResultChan() {
			if w.Object != nil {
				ctx.JSON(w.Object.(models.Event))
			} else {
				ctx.JSON(w.Type)
			}

			ctx.WriteString("\n")
			ctx.ResponseWriter().Flush()
		}
	}
}
