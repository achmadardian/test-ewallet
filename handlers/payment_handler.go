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

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
}

func NewPaymentHandler(p *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: p,
	}
}

func (p *PaymentHandler) Pay(c *gin.Context) {
	var req *requests.PaymentsRequest

	id, ok := helpers.GetUserId(c)
	if !ok {
		log.Printf("failed to get user_id from context")
		responses.InternalServerError(c)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errVal := validate.ExtractValidationErrors(err)
		responses.BadRequest(c, errVal, errs.ErrValidationError.Error())
		return
	}

	tr, payment, err := p.paymentService.Pay(req, id)
	if err != nil {
		if errors.Is(err, errs.ErrInsufficientFund) {
			responses.BadRequest(c, nil, errs.ErrInsufficientFund.Error())
			return
		}

		log.Printf("failed to create payment: %v", err)
		responses.InternalServerError(c)
		return
	}

	res := responses.PaymentResponse{
		Id:             payment.Id,
		Amount:         req.Amount,
		Remarks:        *tr.Remarks,
		BalancedBefore: tr.BalanceBefore,
		BalancedAfter:  tr.BalanceAfter,
		CreatedAt:      tr.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	responses.Ok(c, res)
}
