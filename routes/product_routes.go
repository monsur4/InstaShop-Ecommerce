package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	productGroup := router.Group("/products").Use(middleware.JWTMiddleware())
	{
		productGroup.GET("/", controllers.GetProducts)
		productGroup.GET("/:id", controllers.GetProductByID)
		productGroup.POST("/", controllers.CreateProduct)  // Admin Only
		productGroup.PUT("/:id", controllers.UpdateProduct) // Admin Only
		productGroup.DELETE("/:id", controllers.DeleteProduct) // Admin Only
	}
}

