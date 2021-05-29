package infra

import (
	"eventsmanagergateway/src/entities"
	"eventsmanagergateway/src/utils"

	"github.com/streadway/amqp"
)

var RabbitMqChannel *amqp.Channel

const (
	queueName = "raw_events"
)

type RabbitMqClient struct {
}

func NewRabbitMqClient() *RabbitMqClient {
	return &RabbitMqClient{}
}

func (*RabbitMqClient) Send(event *entities.Event) error {

	q, err := createQueue()

	if err != nil {
		return err
	}

	return RabbitMqChannel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        utils.GenerateBytes(event),
		})
}

func createQueue() (amqp.Queue, error) {
	return RabbitMqChannel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}
