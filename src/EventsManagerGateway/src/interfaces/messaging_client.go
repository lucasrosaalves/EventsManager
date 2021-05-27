package interfaces

import "eventsmanagergateway/src/entities"

type IMessagingClient interface {
	Send(event *entities.Event) error
}
