package routes

import (
	"product-management/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/products", controllers.CreateProduct)
		api.GET("/products/:id", controllers.GetProduct)
	}
}
