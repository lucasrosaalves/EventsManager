package controllers

import (
	"eventsmanagergateway/src/useCases"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	createEventUseCase useCases.ICreateEventUseCase
}

func NewEventController() *EventController {
	return &EventController{
		createEventUseCase: &useCases.CreateEventUseCase{},
	}
}

func (e *EventController) Post(c *gin.Context) {

	var eventRequest struct {
		Timestamp string `json:"timestamp" binding:"required"`
		Tag       string `json:"tag" binding:"required"`
		Value     string `json:"value"`
	}
	c.Bind(&eventRequest)

	status := e.createEventUseCase.Execute(eventRequest.Timestamp, eventRequest.Tag, eventRequest.Value)

	c.JSON(200, gin.H{
		"Status": status,
	})
}
