package usecases

import (
	"eventsmanagervalidator/src/application/interfaces"
	"eventsmanagervalidator/src/domain"
)

type EventsHandler struct {
	messagingService interfaces.MessagingService
}

func NewEventsHandler(messagingService interfaces.MessagingService) domain.EventsHandler {
	return &EventsHandler{
		messagingService: messagingService,
	}
}

func (s *EventsHandler) Execute(eventReceived domain.EventReceived) error {

	event := domain.NewEventValidated(
		eventReceived.Timestamp,
		eventReceived.CountryName,
		eventReceived.RegionName,
		eventReceived.SensorName,
		eventReceived.Value,
		eventReceived.MetaData)

	return s.messagingService.Send(event, event.GetQueueName())
}
