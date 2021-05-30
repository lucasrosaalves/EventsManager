package messaging

import (
	"github.com/streadway/amqp"
)

var RabbitMqContext *rabbitMqContext

type rabbitMqContext struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func CreateRabbitMqContext(url string) error {
	connection, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	RabbitMqContext = &rabbitMqContext{
		Connection: connection,
		Channel:    channel,
	}

	return nil
}
