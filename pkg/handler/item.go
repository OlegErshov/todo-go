package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo_go "todo-go"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo_go.TodoItem
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

type getAllItemsResponse struct {
	Data []todo_go.TodoItem `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo_go.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
