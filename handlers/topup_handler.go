package handlers

import (
	"achmadardian/test-ewallet/requests"
	"achmadardian/test-ewallet/responses"
	"achmadardian/test-ewallet/services"
	"achmadardian/test-ewallet/utils/errs"
	"achmadardian/test-ewallet/utils/helpers"
	"achmadardian/test-ewallet/utils/validate"
	"log"

	"github.com/gin-gonic/gin"
)

type TopupHandler struct {
	topupService *services.TopupService
}

func NewTopupHandler(topupService *services.TopupService) *TopupHandler {
	return &TopupHandler{
		topupService: topupService,
	}
}

func (t *TopupHandler) Topup(c *gin.Context) {
	var req *requests.TopupRequest

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

	tr, topup, err := t.topupService.Create(req, id)
	if err != nil {
		log.Printf("failed to topup: %v", err)
		responses.InternalServerError(c)
		return
	}

	res := responses.TopupResponse{
		Id:             topup.Id,
		AmountTopup:    topup.Amount,
		BalancedBefore: tr.BalanceBefore,
		BalancedAfter:  tr.BalanceAfter,
		CreatedAt:      topup.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	responses.Ok(c, res)
}
