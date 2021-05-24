package routes

import (
	"eventsmanagergateway/src/controllers"

	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) {
	useEventsController(r)
}

func useEventsController(r *gin.Engine) {
	c := controllers.NewEventController()

	r.POST("/events", c.Post)
}
