package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ecommerce-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}

	fmt.Println("✅ Successfully connected to the database")
	DB = database
}

// Apply migrations for all models
func MigrateDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	)
}

