package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Topup struct {
	Id            uuid.UUID      `json:"id" gorm:"primaryKey"`
	TransactionId uuid.UUID      `json:"transaction_id"`
	Amount        int64          `json:"amount"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
