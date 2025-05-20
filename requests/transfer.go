package requests

type TransferRequest struct {
	TargetUser string `json:"target_user" binding:"required"`
	Amount     int64  `json:"amount" binding:"required"`
	Remarks    string `json:"remarks" binding:"required"`
}
