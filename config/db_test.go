package config_test

import (
    "log"
    "os"
    "path/filepath"
    "testing"
    "ecommerce-api/config"
    "github.com/joho/godotenv"
)

func TestConnectDatabase(t *testing.T) {
    // Change the working directory to the project root
    projectRoot, err := filepath.Abs("..")
    if err != nil {
        t.Fatalf("Failed to get project root: %v", err)
    }
    err = os.Chdir(projectRoot)
    if err != nil {
        t.Fatalf("Failed to change working directory: %v", err)
    }

    // Load the .env file
    err = godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Failed to load .env file: %v", err)
    }

    config.ConnectDatabase()

    if config.DB == nil {
        t.Fatal("Database connection is nil")
    }
}