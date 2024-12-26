package models

import "time"

type Product struct {
    ID          uint            `json:"id" gorm:"primaryKey"`
    CreatedAt   time.Time       `json:"created_at"`
    UpdatedAt   time.Time       `json:"updated_at"`
    DeletedAt   *time.Time      `json:"deleted_at" gorm:"index"`
	Name        string          `gorm:"not null"`
	Price       float64         `gorm:"not null"`
	Stock       int             `gorm:"not null"`
}

