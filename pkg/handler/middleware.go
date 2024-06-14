package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		errorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		errorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
	}
	c.Set(userCtx, userId)
}
