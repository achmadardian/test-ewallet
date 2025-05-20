package responses

import "github.com/google/uuid"

type TopupResponse struct {
	Id             uuid.UUID `josn:"id"`
	AmountTopup    int64     `json:"amount_top_up"`
	BalancedBefore int64     `json:"balance_before"`
	BalancedAfter  int64     `json:"balance_after"`
	CreatedAt      string    `json:"created_at"`
}
