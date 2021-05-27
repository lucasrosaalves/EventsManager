package usecases

import (
	"eventsmanagergateway/src/entities"
	"eventsmanagergateway/src/interfaces"
)

type CreateEventRequest struct {
	Timestamp int64  `json:"timestamp" binding:"required"`
	Tag       string `json:"tag" binding:"required"`
	Value     string `json:"value"`
}

type CreateEvent struct {
	messagingService interfaces.IMessagingClient
}

func NewCreateEvent(messagingService interfaces.IMessagingClient) *CreateEvent {
	return &CreateEvent{
		messagingService: messagingService,
	}
}

func (s *CreateEvent) Execute(request *CreateEventRequest) error {
	event, err := entities.NewEvent(request.Timestamp, request.Tag, request.Value)

	if err != nil {
		return err
	}

	return s.messagingService.Send(event)
}
