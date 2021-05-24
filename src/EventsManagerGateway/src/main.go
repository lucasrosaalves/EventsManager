package main

import (
	"eventsmanagergateway/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.UseRoutes(r)
	r.Run(":8080")
}
