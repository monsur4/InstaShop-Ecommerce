package main

import (
	"ecommerce-api/config"
	"ecommerce-api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title InstaShop E-commerce API
// @version 1.0
// @description This is an E-commerce backend API server For InstaShop
// @termsOfService http://swagger.io/terms/
// @contact.name InstaShop Backend
// @contact.url http://www.instashop.com/support
// @contact.email backend@instashop.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Connect to Database
	config.ConnectDatabase()
	// Run AutoMigrate for all models
    config.MigrateDatabase()

	// Initialize Gin router
	router := gin.Default()

	// Test Route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the E-commerce API",
		})
	})

    // Routes
	routes.AuthRoutes(router)
	routes.ProductRoutes(router)
	routes.OrderRoutes(router)

    // Serve Swagger JSON explicitly at a different path
	router.StaticFile("/swagger-docs/doc.json", "./docs/swagger.json")
	// Swagger Documentation Route
    // Serve Swagger UI and explicitly point to doc.json
    router.GET("/swagger/*any", ginSwagger.WrapHandler(
        swaggerFiles.Handler,
        ginSwagger.URL("/swagger-docs/doc.json"),
    ))

	// Start the server
	router.Run(":8080")
}

