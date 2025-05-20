package services

import (
	"achmadardian/test-ewallet/models"
	"achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/utils/pagination"
	"fmt"

	"github.com/google/uuid"
)

type TransactionService struct {
	txRepo *repositories.TransactionRepo
}

func NewTransactionService(txRepo *repositories.TransactionRepo) *TransactionService {
	return &TransactionService{
		txRepo: txRepo,
	}
}

func (t *TransactionService) GetAllByUserId(userId uuid.UUID, p *pagination.Pagination) ([]models.Transaction, error) {
	tr, err := t.txRepo.GetAllByUserId(nil, userId, p)
	if err != nil {
		return nil, fmt.Errorf("get all transaction by id: %w", err)
	}

	return tr, nil
}
