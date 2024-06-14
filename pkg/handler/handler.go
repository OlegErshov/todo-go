package handler

import (
	"github.com/gin-gonic/gin"
	"todo-go/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := router.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := api.Group("/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:id", h.getItemById)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}
	}
	return router
}
