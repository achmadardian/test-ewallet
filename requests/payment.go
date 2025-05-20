package requests

type PaymentsRequest struct {
	Amount  int64  `json:"amount" binding:"required"`
	Remarks string `json:"remarks" binding:"required"`
}
