package responses

import "github.com/google/uuid"

type TransferResponse struct {
	Id            uuid.UUID `json:"id"`
	Status        string    `json:"status"`
	Amount        int64     `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedAt     string    `json:"created_at"`
}
