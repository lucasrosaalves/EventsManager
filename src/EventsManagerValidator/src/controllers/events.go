package controllers

import (
	"encoding/json"
	"eventsmanagervalidator/src/entities"
	"eventsmanagervalidator/src/infra/messaging"
	"log"
)

var (
	queueName string = "event.received"
)

type EventsController struct {
}

func (e *EventsController) post(payload entities.EventReceived) {

}

func NewEventsController() {
	response := make(chan []byte)
	messaging.Consume(queueName, response)

	forever := make(chan bool)

	go func() {
		for d := range response {

			obj := entities.EventReceived{}

			err := json.Unmarshal(d, &obj)

			if err != nil {
				log.Printf(err.Error())
			}
			if err == nil {
				log.Printf("Received a message: %s", d)
			}
		}
	}()

	<-forever
}
