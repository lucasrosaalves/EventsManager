package controllers

import (
	"eventsmanagergateway/src/domain/eventAggregate"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	createEventUseCase eventAggregate.ICreateEventUseCase
}

func NewEventController(CreateEventUseCase eventAggregate.ICreateEventUseCase) *EventController {
	return &EventController{
		createEventUseCase: CreateEventUseCase,
	}
}

func (e *EventController) Post(c *gin.Context) {
	request := &eventAggregate.Event{}

	c.Bind(request)

	status := e.createEventUseCase.Execute(request.Timestamp, request.Tag, request.Value)

	c.JSON(200, gin.H{
		"Status": status,
	})
}
