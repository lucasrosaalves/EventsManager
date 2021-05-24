package main

import (
	"eventsmanagergateway/src/ioc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	ioc.UseRoutes(r)
	r.Run(":8080")
}
