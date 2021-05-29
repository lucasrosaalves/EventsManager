package ioc

import (
	"eventsmanagergateway/src/application/usecases"
	"eventsmanagergateway/src/controllers"
	"eventsmanagergateway/src/infra/messaging"
)

func MakeEventsController() *controllers.EventsController {
	usecase := usecases.NewCreateEvent(messaging.NewRabbitMqClient())

	return controllers.NewEventsController(usecase)
}
