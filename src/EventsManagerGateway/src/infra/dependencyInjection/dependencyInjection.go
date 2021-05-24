package dependencyInjection

import (
	"eventsmanagergateway/src/controllers"
	"eventsmanagergateway/src/useCases"
)

func BuildEventsController() *controllers.EventController {
	return controllers.NewEventController(buildCreateEventUseCase())
}

func buildCreateEventUseCase() *useCases.CreateEventUseCase {
	return &useCases.CreateEventUseCase{}
}
