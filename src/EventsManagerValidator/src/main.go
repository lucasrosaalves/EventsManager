package main

import (
	"encoding/json"
	"eventsmanagervalidator/src/entities"
	"log"

	"github.com/streadway/amqp"
)

var Connection amqp.Connection
var Channel amqp.Channel

func main() {

	useRabbitMq()

	defer Connection.Close()
	defer Channel.Close()
}

func useRabbitMq() {
	Connection, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Channel, _ := Connection.Channel()

	q, _ := Channel.QueueDeclare(
		"event.received", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	msgs, _ := Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			obj := entities.Event{}

			err := json.Unmarshal(d.Body, &obj)

			if err != nil {
				log.Printf(err.Error())
			}
			if err == nil {
				log.Printf("Received a message: %s", d.Body)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
