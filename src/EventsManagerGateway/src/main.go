package main

import (
	"eventsmanagergateway/src/infra/ioc"
	"eventsmanagergateway/src/infra/messaging"

	"github.com/gin-gonic/gin"
)

const (
	port        string = ":8080"
	rabbitMqUrl string = "amqp://guest:guest@localhost:5672/"
)

func main() {
	r := gin.Default()

	useRabbitMq()
	useRoutes(r)

	defer messaging.RabbitMqConnection.Connection.Close()
	defer messaging.RabbitMqConnection.Channel.Close()

	r.Run(port)
}

func useRoutes(r *gin.Engine) {
	ec := ioc.MakeEventsController()
	ec.UseRoutes(r)
}

func useRabbitMq() {
	messaging.CreateRabbitMqConnection(rabbitMqUrl)
}
