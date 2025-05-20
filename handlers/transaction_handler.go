package handlers

import (
	"achmadardian/test-ewallet/responses"
	"achmadardian/test-ewallet/services"
	"achmadardian/test-ewallet/utils/helpers"
	"log"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	txService *services.TransactionService
}

func NewTransactionHandler(txService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		txService: txService,
	}
}

func (t *TransactionHandler) GetAllTransactionByUserId(c *gin.Context) {
	page, err := helpers.GetQueryParamPagination(c)
	if err != nil {
		responses.BadRequest(c, nil, "invalid query params")
		return
	}

	id, ok := helpers.GetUserId(c)
	if !ok {
		log.Printf("failed to get user_id from context")
		responses.InternalServerError(c)
		return
	}

	transactions, err := t.txService.GetAllByUserId(id, page)
	if err != nil {
		log.Printf("failed to fetch all transaction by user id: %v", err)
		responses.InternalServerError(c)
		return
	}

	res := make([]responses.Transaction, 0, len(transactions))
	for _, t := range transactions {
		user := responses.UserTransaction{
			Id:        t.UserId,
			FirstName: t.User.FirstName,
			LastName:  t.User.LastName,
		}

		tx := responses.Transaction{
			Id:              t.Id,
			Status:          t.Status,
			UserId:          t.UserId,
			TransactionType: t.TransactionType,
			Amount:          t.Amount,
			Remarks:         t.Remarks,
			BalanceBefore:   t.BalanceBefore,
			BalanceAfter:    t.BalanceAfter,
			CreatedAt:       t.CreatedAt.Format("2006-01-02 15:04:05"),
			User:            &user,
		}

		res = append(res, tx)
	}

	responses.OkPage(c, res, page)
}
