package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type errorResp struct {
	Message string
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	log.Default().Print(message)
	c.AbortWithStatusJSON(statusCode, errorResp{message})
}

type statusResponse struct {
	Status string `json:"status"`
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		errorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	return id.(int), nil
}
