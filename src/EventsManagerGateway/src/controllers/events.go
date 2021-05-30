package controllers

import (
	"eventsmanagergateway/src/domain"

	"github.com/gin-gonic/gin"
)

type EventsController struct {
	createEventUseCase domain.EventsHandler
}

func NewEventsController(r *gin.Engine, createEventUseCase domain.EventsHandler) {
	controller := &EventsController{
		createEventUseCase: createEventUseCase,
	}

	r.POST("/events", controller.post)
}

func (e *EventsController) post(c *gin.Context) {
	type request struct {
		Timestamp int64  `json:"timestamp" binding:"required"`
		Tag       string `json:"tag" binding:"required"`
		Value     string `json:"value"`
	}

	payload := &request{}
	err := BindJson(c, payload)

	if err != nil {
		BadRequest(c, err)
	}

	err = e.createEventUseCase.Execute(payload.Timestamp, payload.Tag, payload.Value)

	if err != nil {
		BadRequest(c, err)
	} else {
		Ok(c)
	}
}
