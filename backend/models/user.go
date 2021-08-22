package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  []byte         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
