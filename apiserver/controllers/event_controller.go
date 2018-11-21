package controllers

import (
	"github.com/joshuakwan/hydra/apiserver/services"
	"github.com/joshuakwan/hydra/models"
	"github.com/kataras/iris"
)

// EventController handles HTTP requests on /events
type EventController struct {
	Service services.EventService
}

// Get handles GET /events
func (e *EventController) Get() (results []*models.Event) {
	return e.Service.List()
}

// Post handles POST /events
func (e *EventController) Post(ctx iris.Context) error {
	var event models.Event
	err := ctx.ReadJSON(&event)
	if err != nil {
		return err
	}
	return e.Service.Create(&event)
}

// GetWatch is the watch interface for events at /events/watch
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
