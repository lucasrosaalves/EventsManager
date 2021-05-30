package controllers

import (
	"encoding/json"
	"eventsmanagervalidator/src/application/interfaces"
	"eventsmanagervalidator/src/domain"
	"log"
)

var (
	queueName string = "event.received"
)

func NewEventsController(eventsHandler domain.EventsHandler, messagingService interfaces.MessagingService) {
	msgs := make(chan []byte)

	go messagingService.Subscribe(queueName, msgs)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d)

			payload := domain.EventReceived{}

			err := json.Unmarshal(d, &payload)
			if err == nil {
				eventsHandler.Execute(payload)
			}
		}
	}()

	<-forever
}
