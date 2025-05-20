package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	Id              uuid.UUID      `json:"id" gorm:"primaryKey"`
	Status          string         `json:"status"`
	UserId          uuid.UUID      `json:"user_id"`
	TransactionType string         `json:"transaction_type"`
	Amount          int64          `json:"amount"`
	Remarks         *string        `json:"remarks"`
	BalanceBefore   int64          `json:"balance_before"`
	BalanceAfter    int64          `json:"balance_after"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	User            *User          `json:"user" gorm:"foreignKey:UserId;references:Id"`
}
