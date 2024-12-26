package controllers

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PlaceOrder godoc
// @Summary Place an order
// @Description Place an order for products
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body models.Order true "Order Details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /orders [post]
func PlaceOrder(c *gin.Context) {
	userID := c.GetUint("userId")
    if userID == 0 {
    	c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
    	return
    }

	var orderInput struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, orderInput.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.Stock < orderInput.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	order := models.Order{
		UserID:    userID,
		ProductID: orderInput.ProductID,
		Quantity:  orderInput.Quantity,
		Status:    "Pending",
		OrderDate: time.Now(),
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	// Update Product Stock
	product.Stock -= orderInput.Quantity
	config.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order": order})
}

// ListUserOrders godoc
// @Summary List all orders for a user
// @Description Retrieve all orders made by the authenticated user
// @Tags Order
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Order
// @Failure 401 {object} map[string]string
// @Router /orders [get]
func ListUserOrders(c *gin.Context) {
	userID := c.GetUint("userId")

	var orders []models.Order
	config.DB.Where("user_id = ?", userID).Find(&orders)

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// CancelOrder godoc
// @Summary Cancel an order
// @Description Cancel an order if it is still pending
// @Tags Order
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /orders/{id}/cancel [put]
func CancelOrder(c *gin.Context) {
	userID := c.GetUint("userId")
	orderID := c.Param("id")

	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if order.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only cancel your own orders"})
		return
	}

	if order.Status != "Pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only pending orders can be cancelled"})
		return
	}

	order.Status = "Cancelled"
	config.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

// UpdateOrderStatus godoc
// @Summary Update the status of an order
// @Description Update an order's status (Admin only)
// @Tags Order
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body map[string]string true "Order Status (Pending, Completed, Cancelled)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /orders/{id}/status [put]
func UpdateOrderStatus(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	orderID := c.Param("id")
	var order models.Order

	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var statusInput struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&statusInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.Status = statusInput.Status
	config.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully", "order": order})
}

