package controllers

import "github.com/gin-gonic/gin"

func BindJson(c *gin.Context, obj interface{}) error {
	return c.ShouldBindJSON(obj)
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"Message": err.Error(),
	})
}

func Ok(c *gin.Context) {
	c.Status(200)
}
