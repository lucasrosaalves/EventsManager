package controllers

import (
	"eventsmanagergateway/src/interfaces"
	"eventsmanagergateway/src/models"

	"github.com/gin-gonic/gin"
)

type EventsController struct {
	eventsService interfaces.IEventsService
}

func NewEventsController(eventsService interfaces.IEventsService) *EventsController {
	return &EventsController{
		eventsService: eventsService,
	}
}

func (e *EventsController) Post(c *gin.Context) {
	request := &models.Event{}

	c.Bind(request)

	status := e.eventsService.Execute(request.Timestamp, request.Tag, request.Value)

	c.JSON(200, gin.H{
		"Status": status,
	})
}
