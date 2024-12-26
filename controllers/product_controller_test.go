package controllers

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	config.DB.Create(&models.Product{Name: "Test Product", Price: 10.0, Stock: 5})

	router := gin.Default()
	router.GET("/products", GetProducts)

	token := loginAndGetToken(t)

	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Product")
}

