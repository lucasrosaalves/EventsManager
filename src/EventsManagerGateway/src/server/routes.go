package server

import (
	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) {
	useEventsRoutes(r)
}
