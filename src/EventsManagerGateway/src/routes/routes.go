package routes

import (
	"eventsmanagergateway/src/infra/dependencyInjection"

	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) {
	useEventsController(r)
}

func useEventsController(r *gin.Engine) {
	c := dependencyInjection.BuildEventsController()

	r.POST("/events", c.Post)
}
