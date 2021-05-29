package usecases

import (
	"eventsmanagergateway/src/application/interfaces"
	"eventsmanagergateway/src/application/utils"
	"eventsmanagergateway/src/domain"
)

type CreateEvent struct {
	messagingService interfaces.MessagingService
}

func NewCreateEvent(messagingService interfaces.MessagingService) domain.CreateEventUseCase {
	return &CreateEvent{
		messagingService: messagingService,
	}
}

func (s *CreateEvent) Execute(timestamp int64, tag string, value string) error {
	correlationId := utils.GenerateGuid()

	event, err := domain.NewEvent(timestamp, tag, value, correlationId)

	if err != nil {
		return err
	}

	return s.messagingService.Send(event, event.GetQueueName())
}
