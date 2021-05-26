package main

import (
	"eventsmanagergateway/src/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	server.UseRoutes(r)
	r.Run(":8080")
}
