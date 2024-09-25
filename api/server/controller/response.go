package controller

import "github.com/gin-gonic/gin"

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {

	c.AbortWithStatusJSON(statusCode, Error{
		Message: message,
	})
}
