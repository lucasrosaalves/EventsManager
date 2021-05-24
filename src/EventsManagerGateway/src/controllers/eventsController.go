package controllers

import (
	"eventsmanagergateway/src/domain/eventAggregate"

	"github.com/gin-gonic/gin"
)

type EventsController struct {
	createEventUseCase eventAggregate.ICreateEventUseCase
}

func NewEventsController(CreateEventUseCase eventAggregate.ICreateEventUseCase) *EventsController {
	return &EventsController{
		createEventUseCase: CreateEventUseCase,
	}
}

func (e *EventsController) Post(c *gin.Context) {
	request := &eventAggregate.Event{}

	c.Bind(request)

	status := e.createEventUseCase.Execute(request.Timestamp, request.Tag, request.Value)

	c.JSON(200, gin.H{
		"Status": status,
	})
}
