package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type error struct {
	Message string
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	log.Default().Print(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
