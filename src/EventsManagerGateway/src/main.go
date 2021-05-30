package main

import (
	"eventsmanagergateway/src/application/usecases"
	"eventsmanagergateway/src/controllers"
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
	controllers.NewEventsController(r, usecases.NewCreateEvent(messaging.NewRabbitMqClient()))
}

func useRabbitMq() {
	messaging.CreateRabbitMqConnection(rabbitMqUrl)
}
