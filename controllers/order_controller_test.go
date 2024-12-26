package controllers

import (
    "bytes"
    "ecommerce-api/config"
    "ecommerce-api/middleware"
    "ecommerce-api/models"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestPlaceOrder(t *testing.T) {
    // Mock data
    user := models.User{Email: "test@example.com", Password: "password"}
    product := models.Product{Name: "Test Product", Price: 10.0, Stock: 5}
    config.DB.Create(&user)
    config.DB.Create(&product)

    token := loginAndGetToken(t)

    router := gin.Default()
    router.Use(middleware.JWTMiddleware())
    router.POST("/orders", PlaceOrder)

    orderData := map[string]interface{}{
        "product_id": product.ID,
        "quantity":   1,
    }

    jsonValue, _ := json.Marshal(orderData)
    req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Order placed successfully")
}