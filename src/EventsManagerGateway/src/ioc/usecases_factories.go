package ioc

import (
	"eventsmanagergateway/src/infra"
	"eventsmanagergateway/src/usecases"
)

func BuildCreateEvent() *usecases.CreateEvent {
	return usecases.NewCreateEvent(infra.NewRabbitMqClient())
}
