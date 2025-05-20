package responses

import (
	"github.com/google/uuid"
)

type UserTransaction struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type Transaction struct {
	Id              uuid.UUID        `json:"id" gorm:"primaryKey"`
	Status          string           `json:"status"`
	UserId          uuid.UUID        `json:"user_id"`
	TransactionType string           `json:"transaction_type"`
	Amount          int64            `json:"amount"`
	Remarks         *string          `json:"remarks"`
	BalanceBefore   int64            `json:"balance_before"`
	BalanceAfter    int64            `json:"balance_after"`
	CreatedAt       string           `json:"created_at"`
	User            *UserTransaction `json:"user"`
}
