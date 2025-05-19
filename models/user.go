package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id          uuid.UUID      `json:"id" gorm:"primaryKey"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	PhoneNumber string         `json:"phone_number"`
	Address     string         `json:"address"`
	Pin         string         `json:"pin"`
	Balance     int64          `json:"balance"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
