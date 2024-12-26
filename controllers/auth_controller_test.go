package controllers

import (
    "bytes"
    "ecommerce-api/config"
    "ecommerce-api/models"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
    // Cleanup: delete any existing user with the same email
    config.DB.Where("email = ?", "test@example.com").Delete(&models.User{})

    router := gin.Default()
    router.POST("/auth/register", Register)

    // Test Data
    user := map[string]string{
        "email":    "test@example.com",
        "password": "password123",
    }

    jsonValue, _ := json.Marshal(user)
    req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "User registered successfully")
}

func loginAndGetToken(t *testing.T) string {
    router := gin.Default()
    router.POST("/auth/login", Login)

    // Login Data
    loginData := map[string]string{
        "email":    "test@example.com",
        "password": "password123",
    }

    jsonValue, _ := json.Marshal(loginData)
    req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    if err != nil {
        t.Fatalf("Failed to parse login response: %v", err)
    }

    token, exists := response["token"]
    if !exists {
        t.Fatalf("Token not found in login response")
    }

    return token
}