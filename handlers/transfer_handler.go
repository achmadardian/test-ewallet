package handlers

import (
	"achmadardian/test-ewallet/requests"
	"achmadardian/test-ewallet/responses"
	"achmadardian/test-ewallet/services"
	"achmadardian/test-ewallet/utils/errs"
	"achmadardian/test-ewallet/utils/helpers"
	"achmadardian/test-ewallet/utils/validate"
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type TransferHandler struct {
	tr *services.TransferService
}

func NewTransferHandler(tr *services.TransferService) *TransferHandler {
	return &TransferHandler{
		tr: tr,
	}
}

func (t *TransferHandler) Transfer(c *gin.Context) {
	var req *requests.TransferRequest

	id, ok := helpers.GetUserId(c)
	if !ok {
		log.Printf("failed to get user_id from context")
		responses.InternalServerError(c)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errRes := validate.ExtractValidationErrors(err)
		responses.BadRequest(c, errRes, errs.ErrValidationError.Error())
		return
	}

	pay, err := t.tr.Transfer(req, id)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			responses.NotFound(c, errs.ErrDataNotFound.Error())
			return
		}

		if errors.Is(err, errs.ErrInsufficientFund) {
			responses.BadRequest(c, nil, errs.ErrInsufficientFund.Error())
			return
		}

		log.Printf("failed to transfer: %v", err)
		responses.InternalServerError(c)
		return
	}

	res := responses.TransferResponse{
		Id:            pay.Id,
		Status:        "PENDING",
		Amount:        req.Amount,
		Remarks:       req.Remarks,
		BalanceBefore: pay.SenderBalanceBefore,
		BalanceAfter:  pay.SenderBalanceAfter,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	responses.Ok(c, res)
}
