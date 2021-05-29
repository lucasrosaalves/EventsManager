package main

import (
	"eventsmanagergateway/src/controllers"
	"eventsmanagergateway/src/infra"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

var Connection amqp.Connection
var Channel amqp.Channel

func main() {
	r := gin.Default()

	useRabbitMq()
	useRoutes(r)

	defer Connection.Close()
	defer Channel.Close()

	r.Run(":8080")
}

func useRabbitMq() {
	Connection, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Channel, _ := Connection.Channel()

	infra.RabbitMqChannel = Channel
}

func useRoutes(r *gin.Engine) {
	r.POST("/events", controllers.PostEvent)
}
