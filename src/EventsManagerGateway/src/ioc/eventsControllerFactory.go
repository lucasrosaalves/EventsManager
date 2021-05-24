package ioc

import (
	"eventsmanagergateway/src/controllers"
	"eventsmanagergateway/src/useCases"
)

func buildEventsController() *controllers.EventsController {
	return controllers.NewEventsController(&useCases.CreateEventUseCase{})
}
