package main

import (
	"eventsmanagergateway/src/infra/ioc"
	"eventsmanagergateway/src/infra/messaging"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	r := gin.Default()

	useRabbitMq()
	useRoutes(r)

	defer messaging.Connection.Close()
	defer messaging.Channel.Close()

	r.Run(":8080")
}

func useRoutes(r *gin.Engine) {
	ec := ioc.MakeEventsController()
	ec.UseRoutes(r)
}

func useRabbitMq() {
	messaging.Connection, _ = amqp.Dial("amqp://guest:guest@localhost:5672/")
	messaging.Channel, _ = messaging.Connection.Channel()
}
