package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transfer struct {
	Id                    uuid.UUID      `json:"id" gorm:"primaryKey"`
	SenderTransactionId   uuid.UUID      `json:"sender_transaction_id"`
	ReceiverTransactionId uuid.UUID      `json:"receiver_transaction_id"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}
