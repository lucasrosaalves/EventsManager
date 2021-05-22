package main

import (
	"eventsmanagergateway/src/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {

		result := controllers.HandleEvent()
		c.JSON(200, gin.H{
			"Status": result,
		})
	})
}

func handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	setupRouter(r)
	r.Run(":8080")
}
