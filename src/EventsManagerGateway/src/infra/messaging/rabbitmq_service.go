package messaging

import (
	"eventsmanagergateway/src/application/utils"

	"github.com/streadway/amqp"
)

const (
	queueName = "raw_events"
)

type RabbitMqService struct {
}

func NewRabbitMqClient() *RabbitMqService {
	return &RabbitMqService{}
}

func (*RabbitMqService) Send(obj interface{}) error {

	q, err := createQueue()

	if err != nil {
		return err
	}

	return RabbitMqConnection.Channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        utils.GenerateBytes(obj),
		})
}

func createQueue() (amqp.Queue, error) {
	return RabbitMqConnection.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}
