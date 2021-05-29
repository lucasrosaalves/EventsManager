package controllers

import (
	response "eventsmanagergateway/src/application/responses"
	"eventsmanagergateway/src/domain"

	"github.com/gin-gonic/gin"
)

func (e *EventsController) UseRoutes(r *gin.Engine) {
	r.POST("/events", e.post)
}

type EventsController struct {
	createEventUseCase domain.CreateEventUseCase
}

func NewEventsController(createEventUseCase domain.CreateEventUseCase) *EventsController {
	return &EventsController{
		createEventUseCase: createEventUseCase,
	}
}

func (e *EventsController) post(c *gin.Context) {
	type request struct {
		Timestamp int64  `json:"timestamp" binding:"required"`
		Tag       string `json:"tag" binding:"required"`
		Value     string `json:"value"`
	}

	payload := &request{}
	err := c.ShouldBindJSON(payload)

	err = e.createEventUseCase.Execute(payload.Timestamp, payload.Tag, payload.Value)

	if err != nil {
		c.JSON(400, response.NewBadRequest(err))
	} else {
		c.Status(202)
	}

}
