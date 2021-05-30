package main

import (
	"eventsmanagervalidator/src/application/usecases"
	"eventsmanagervalidator/src/controllers"
	"eventsmanagervalidator/src/infra/messaging"

	"github.com/streadway/amqp"
)

const (
	rabbitMqUrl string = "amqp://guest:guest@localhost:5672/"
)

var Connection amqp.Connection
var Channel amqp.Channel

func main() {
	messaging.CreateRabbitMqConnection(rabbitMqUrl)
	messagingService := messaging.NewRabbitMqService()

	controllers.NewEventsController(
		usecases.NewEventsHandler(messagingService),
		messagingService)

	defer messaging.RabbitMqConnection.Connection.Close()
	defer messaging.RabbitMqConnection.Channel.Close()
}
