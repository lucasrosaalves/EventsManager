package controllers

import "github.com/gin-gonic/gin"

func BadRequest(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"Message": err.Error(),
	})
}

func Ok(c *gin.Context) {
	c.Status(200)
}
