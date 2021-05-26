package server

import (
	"eventsmanagergateway/src/controllers"
	"eventsmanagergateway/src/services"

	"github.com/gin-gonic/gin"
)

func useEventsRoutes(r *gin.Engine) {
	c := buildEventsController()

	r.POST("/events", c.Post)
}

func buildEventsController() *controllers.EventsController {
	return controllers.NewEventsController(&services.EventsService{})
}
