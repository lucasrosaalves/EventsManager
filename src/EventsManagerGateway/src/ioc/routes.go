package ioc

import (
	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) {
	useEventsController(r)
}

func useEventsController(r *gin.Engine) {
	c := buildEventsController()

	r.POST("/events", c.Post)
}
