package controllers

import (
	"encoding/json"
	"eventsmanagervalidator/src/application/interfaces"
	"eventsmanagervalidator/src/entities"
	"log"
)

var (
	queueName string = "event.received"
)

type EventsController struct {
}

func (e *EventsController) post(payload entities.EventReceived) {

}

func NewEventsController(messagingService interfaces.MessagingService) {
	msgs := make(chan []byte)

	go messagingService.Subscribe(queueName, msgs)

	controller := &EventsController{}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d)

			payload := entities.EventReceived{}

			err := json.Unmarshal(d, &payload)
			if err == nil {
				controller.post(payload)
			}
		}
	}()

	<-forever
}
