package responses

import "github.com/google/uuid"

type PaymentResponse struct {
	Id             uuid.UUID `json:"id"`
	Amount         int64     `json:"amount"`
	Remarks        string    `json:"remarks"`
	BalancedBefore int64     `json:"balance_before"`
	BalancedAfter  int64     `json:"balance_after"`
	CreatedAt      string    `json:"created_at"`
}
