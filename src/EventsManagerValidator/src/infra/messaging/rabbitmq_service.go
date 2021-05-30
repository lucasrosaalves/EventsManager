package messaging

import (
	"eventsmanagervalidator/src/application/interfaces"
	"eventsmanagervalidator/src/application/utils"

	"github.com/streadway/amqp"
)

type RabbitMqService struct {
}

func NewRabbitMqService() interfaces.MessagingService {
	return &RabbitMqService{}
}

func (*RabbitMqService) Send(obj interface{}, queueName string) error {
	queue, err := createQueue(queueName)

	if err != nil {
		return err
	}

	publishObj := createObject(obj)
	return publish(publishObj, queue, queueName)
}

func (*RabbitMqService) Subscribe(queueName string, c chan []byte) {
	channel := RabbitMqConnection.Channel

	q, _ := createQueue(queueName)

	msgs, _ := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for d := range msgs {
		c <- d.Body
	}
}

func publish(publishObj amqp.Publishing, queue amqp.Queue, queueName string) error {
	return RabbitMqConnection.Channel.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		publishObj)
}

func createObject(obj interface{}) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "text/plain",
		Body:        utils.GenerateBytes(obj),
	}
}

func createQueue(queueName string) (amqp.Queue, error) {
	return RabbitMqConnection.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}
