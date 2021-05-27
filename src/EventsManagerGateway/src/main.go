package main

import (
	"eventsmanagergateway/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	useRoutes(r)
	r.Run(":8080")
}

func useRoutes(r *gin.Engine) {
	r.POST("/events", controllers.PostEvent)
}
