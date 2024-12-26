package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	orderGroup := router.Group("/orders").Use(middleware.JWTMiddleware())
	{
		orderGroup.POST("/", controllers.PlaceOrder)
		orderGroup.GET("/", controllers.ListUserOrders)
		orderGroup.PUT("/:id/cancel", controllers.CancelOrder)
		orderGroup.PUT("/:id/status", controllers.UpdateOrderStatus) // Admin Only
	}
}

