package infra

import "eventsmanagergateway/src/entities"

type RabbitMqClient struct {
}

func NewRabbitMqClient() *RabbitMqClient {
	return &RabbitMqClient{}
}

func (*RabbitMqClient) Send(event *entities.Event) error {

	return nil
}
