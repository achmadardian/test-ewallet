package services

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"
	repo "achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/requests"
	"achmadardian/test-ewallet/utils/errs"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentService struct {
	paymentRepo *repo.PaymentRepo
	userService *UserService
	txService   *TransactionService
	DB          db.Database
}

func NewPaymentService(paymentRepo *repo.PaymentRepo, userService *UserService, txService *TransactionService, DB db.Database) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
		userService: userService,
		txService:   txService,
		DB:          DB,
	}
}

func (p *PaymentService) Pay(req *requests.PaymentsRequest, userId uuid.UUID) (*models.Transaction, *models.Payment, error) {
	// tx
	tx := p.DB.Write().Begin()

	// get balance
	balance, err := p.userService.GetBalance(tx, userId)
	if err != nil {
		return nil, nil, fmt.Errorf("get balance for payment: %w", err)
	}

	oldBalance := balance.Balance
	newBalance := oldBalance - req.Amount
	status := repo.FilterStatusSuccess
	trType := repo.FilterTransactionDebit

	// check balance sufficient
	if req.Amount > oldBalance {
		tx.Rollback()
		return nil, nil, errs.ErrInsufficientFund
	}

	// create transaction
	createTr, err := p.txService.Create(tx, status, userId, trType, req.Amount, oldBalance, newBalance, &req.Remarks)
	if err != nil {
		return nil, nil, fmt.Errorf("create transaction for payment: %w", err)
	}

	// create payment
	createPayment, err := p.create(tx, createTr.Id)
	if err != nil {
		return nil, nil, fmt.Errorf("create payment: %w", err)
	}

	// update balance
	_, err = p.userService.UpdateBalance(tx, userId, newBalance)
	if err != nil {
		return nil, nil, fmt.Errorf("update balance for payment: %w", err)
	}

	tx.Commit()

	return createTr, createPayment, nil
}

func (p *PaymentService) create(tx *gorm.DB, transactionId uuid.UUID) (*models.Payment, error) {
	payment := &models.Payment{
		Id:            uuid.New(),
		TransactionId: transactionId,
	}

	create, err := p.paymentRepo.Create(tx, payment)
	if err != nil {
		return nil, err
	}

	return create, nil
}
