package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	todo_go "todo-go"
)

func (h *Handler) SignIn(c *gin.Context) {

}

func (h *Handler) SignUp(c *gin.Context) {
	var input todo_go.User

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
