package requests

type TopupRequest struct {
	Amount int64 `json:"amount" binding:"required"`
}
