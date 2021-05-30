package usecases

import (
	"eventsmanagergateway/src/application/interfaces"
	"eventsmanagergateway/src/application/utils"
	"eventsmanagergateway/src/domain"
)

type EventsHandler struct {
	messagingService interfaces.MessagingService
}

func NewEventsHandler(messagingService interfaces.MessagingService) domain.EventsHandler {
	return &EventsHandler{
		messagingService: messagingService,
	}
}

func (s *EventsHandler) Execute(timestamp int64, tag string, value string) error {
	correlationId := utils.GenerateGuid()

	event, err := domain.NewEvent(timestamp, tag, value, correlationId)

	if err != nil {
		return err
	}

	return s.messagingService.Send(event, event.GetQueueName())
}
