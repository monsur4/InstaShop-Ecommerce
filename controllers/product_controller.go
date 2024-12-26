package controllers

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Create a product
// @Description Create a new product (Admin only)
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body models.Product true "Product Details"
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}

// GetProducts godoc
// @Summary List all products
// @Description Retrieve all products
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Retrieve a product by its unique ID
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product's details (Admin only)
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated Product Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": product})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Remove a product from the catalog (Admin only)
// @Tags Product
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	id := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

