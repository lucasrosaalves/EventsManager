package controllers

import (
	"eventsmanagergateway/src/ioc"
	"eventsmanagergateway/src/usecases"

	"github.com/gin-gonic/gin"
)

func PostEvent(c *gin.Context) {
	request := &usecases.CreateEventRequest{}
	err := c.ShouldBindJSON(request)

	useCase := ioc.BuildCreateEvent()

	err = useCase.Execute(request)

	if err != nil {
		c.JSON(400, gin.H{
			"Message": err.Error(),
		})
	} else {
		c.Status(202)
	}

}
