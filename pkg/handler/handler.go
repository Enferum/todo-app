package handler

import (
	"github.com/Enferum/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		api.POST("/", h.createList)
		api.GET("/", h.getAllLists)
		api.GET("/:id", h.getListById)
		api.PUT("/:id", h.updateList)
		api.DELETE("/:id", h.deleteList)

		items := router.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
