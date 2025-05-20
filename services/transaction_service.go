package services

import (
	"achmadardian/test-ewallet/models"
	repo "achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/utils/pagination"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionService struct {
	txRepo *repo.TransactionRepo
}

func NewTransactionService(txRepo *repo.TransactionRepo) *TransactionService {
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

func (t *TransactionService) Create(tx *gorm.DB, status repo.FilterValue, userId uuid.UUID, trType repo.FilterValue, amount, balBefore, balAfter int64) (*models.Transaction, error) {
	tr := &models.Transaction{
		Id:              uuid.New(),
		Status:          string(status),
		UserId:          userId,
		TransactionType: string(trType),
		Amount:          amount,
		BalanceBefore:   balBefore,
		BalanceAfter:    balAfter,
	}

	save, err := t.txRepo.Create(tx, tr)
	if err != nil {
		return nil, err
	}

	return save, err
}
