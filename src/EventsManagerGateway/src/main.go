package main

import (
	"eventsmanagergateway/src/controllers"
	"eventsmanagergateway/src/infra"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	r := gin.Default()

	useRabbitMq()
	useRoutes(r)

	defer infra.Connection.Close()
	defer infra.Channel.Close()

	r.Run(":8080")
}

func useRoutes(r *gin.Engine) {
	r.POST("/events", controllers.PostEvent)
}

func useRabbitMq() {
	infra.Connection, _ = amqp.Dial("amqp://guest:guest@localhost:5672/")
	infra.Channel, _ = infra.Connection.Channel()
}
