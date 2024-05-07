package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

// инициализирует все ендпоинты
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up")
		auth.POST("/sing-in")
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")
		}

		items := lists.Group("/")
		{
			items.POST("/")
			items.GET("/")
			items.GET("/:item_id")
			items.PUT("/:item_id")
			items.DELETE("/:item_id")
		}
	}

	return router
}