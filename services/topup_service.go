package services

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"
	repo "achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/requests"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TopupService struct {
	topupRepo   *repo.TopupRepo
	userService *UserService
	txService   *TransactionService
	DB          db.Database
}

func NewTopService(
	topupRepo *repo.TopupRepo,
	userService *UserService,
	txService *TransactionService,
	DB db.Database) *TopupService {
	return &TopupService{
		topupRepo:   topupRepo,
		userService: userService,
		txService:   txService,
		DB:          DB,
	}
}

func (t *TopupService) Create(req *requests.TopupRequest, userId uuid.UUID) (*models.Transaction, *models.Topup, error) {
	// transaction
	tx := t.DB.Write().Begin()

	// get balance
	balance, err := t.userService.GetBalance(tx, userId)
	if err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("get user balance for update: %w", err)
	}

	oldBalance := balance.Balance
	newBalance := balance.Balance + req.Amount
	status := repo.FilterStatusSuccess
	trType := repo.FilterTransactionCredit

	// create transaction
	createTr, err := t.txService.Create(tx, status, userId, trType, req.Amount, oldBalance, newBalance, nil)
	if err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("create transaction for topup: %w", err)
	}

	// create topup
	createTopup, err := t.create(tx, createTr.Id, req.Amount)
	if err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("create topup: %w", err)
	}

	// update balance
	_, err = t.userService.UpdateBalance(tx, userId, newBalance)
	if err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("update user balance: %w", err)
	}

	// commit
	tx.Commit()

	return createTr, createTopup, nil
}

func (t *TopupService) create(tx *gorm.DB, transactionId uuid.UUID, amount int64) (*models.Topup, error) {
	topup := &models.Topup{
		Id:            uuid.New(),
		TransactionId: transactionId,
		Amount:        amount,
	}
	create, err := t.topupRepo.Create(tx, topup)
	if err != nil {
		return nil, err
	}

	return create, nil
}
