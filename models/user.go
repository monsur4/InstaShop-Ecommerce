package models

import "time"

type User struct {
    ID          uint            `json:"id" gorm:"primaryKey"`
    CreatedAt   time.Time       `json:"created_at"`
    UpdatedAt   time.Time       `json:"updated_at"`
    DeletedAt   *time.Time      `json:"deleted_at" gorm:"index"`
	Email       string          `gorm:"unique;not null"`
	Password    string          `gorm:"not null"`
	Role        string          `gorm:"default:user"` // Can be 'user' or 'admin'
}

