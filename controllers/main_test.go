package controllers

import (
    "ecommerce-api/config"
    "ecommerce-api/models"
    "os"
    "path/filepath"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
    // Change the working directory to the project root
    projectRoot, err := filepath.Abs("..")
    if err != nil {
        panic("Failed to get project root: " + err.Error())
    }
    err = os.Chdir(projectRoot)
    if err != nil {
        panic("Failed to change working directory: " + err.Error())
    }

    // Load the .env file
    err = godotenv.Load(".env")
    if err != nil {
        panic("Failed to load .env file: " + err.Error())
    }

    // Set test-specific environment variables
    os.Setenv("DB_NAME", "ecommerce_test")

    // Setup
    gin.SetMode(gin.TestMode)
    config.ConnectDatabase()
    config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

    // Run tests
    code := m.Run()

    os.Exit(code)
}