package models

import "time"

type Order struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt *time.Time        `json:"deleted_at" gorm:"index"`
	UserID    uint              `gorm:"not null"`
	ProductID uint              `gorm:"not null"`
	Quantity  int               `gorm:"not null"`
	Status    string            `gorm:"default:Pending"` // Possible values: Pending, Completed, Cancelled
	OrderDate time.Time         `gorm:"autoCreateTime"`
}

