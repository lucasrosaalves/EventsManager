package messaging

import (
	"github.com/streadway/amqp"
)

var RabbitMqConnection *rabbitMqConnection

type rabbitMqConnection struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func CreateRabbitMqConnection(url string) error {
	connection, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	RabbitMqConnection = &rabbitMqConnection{
		Connection: connection,
		Channel:    channel,
	}

	return nil
}
