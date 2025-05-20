package services

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"
	"achmadardian/test-ewallet/queue"
	repo "achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/requests"
	"achmadardian/test-ewallet/utils/errs"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransferService struct {
	transferRepo *repo.TransferRepo
	userService  *UserService
	rabbitmq     *queue.RabbitMqClient
	txService    *TransactionService
	DB           db.Database
}

type TransferInfo struct {
	TransferId    uuid.UUID
	BalanceBefore int64
	BalanceAfter  int64
}

func NewTransferService(trasnferRepo *repo.TransferRepo, userService *UserService, rabbitmq *queue.RabbitMqClient, DB db.Database, txService *TransactionService) *TransferService {
	return &TransferService{
		transferRepo: trasnferRepo,
		userService:  userService,
		rabbitmq:     rabbitmq,
		DB:           DB,
		txService:    txService,
	}
}

func (t *TransferService) Transfer(req *requests.TransferRequest, senderId uuid.UUID) (*queue.Transfer, error) {
	receiverId, err := uuid.Parse(req.TargetUser)
	if err != nil {
		return &queue.Transfer{}, errs.ErrDataNotFound
	}

	senderBalance, err := t.userService.GetAccount(senderId)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			return &queue.Transfer{}, errs.ErrDataNotFound
		}

		return &queue.Transfer{}, fmt.Errorf("get sender account for transfer: %w", err)
	}
	receiver, err := t.userService.GetAccount(receiverId)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			return &queue.Transfer{}, errs.ErrDataNotFound
		}

		return &queue.Transfer{}, fmt.Errorf("get receiver account for transfer: %w", err)
	}

	oldSenderBalance := senderBalance.Balance
	newSenderBalance := oldSenderBalance - req.Amount
	oldReceiverBalance := receiver.Balance
	newReceiverBalance := oldReceiverBalance + req.Amount
	if req.Amount > oldSenderBalance {
		return &queue.Transfer{}, errs.ErrInsufficientFund
	}

	// publish queue
	tfId := uuid.New()
	msg := queue.Transfer{
		Id:                     tfId,
		SenderId:               senderId,
		ReceiverId:             receiverId,
		SenderBalanceBefore:    oldSenderBalance,
		SenderBalanceAfter:     newSenderBalance,
		ReceiverBalancerBefore: oldReceiverBalance,
		ReceiverBalanceAfter:   newReceiverBalance,
		Amount:                 req.Amount,
		Remarks:                req.Remarks,
	}
	if err := t.rabbitmq.PublishTransfer(msg, "transfer"); err != nil {
		return nil, fmt.Errorf("send message for transfer: %w", err)
	}

	return &queue.Transfer{
		Id:                  tfId,
		SenderBalanceBefore: oldSenderBalance,
		SenderBalanceAfter:  newSenderBalance,
	}, nil
}

func (t *TransferService) HandleTransfer(msg *queue.Transfer) error {
	// tx
	tx := t.DB.Write().Begin()

	// recalculate balance sender
	_, err := t.userService.GetBalance(tx, msg.SenderId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("get sender balance for handle transfer: %w", err)
	}
	_, err = t.userService.GetBalance(tx, msg.ReceiverId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("get receiver balance for handle transfer: %w", err)
	}

	// insert into table transaction for sender
	sendStatus := repo.FilterStatusSuccess
	sendTxType := repo.FilterTransactionDebit
	sendTx, err := t.txService.Create(tx, sendStatus, msg.SenderId, sendTxType, msg.Amount, msg.SenderBalanceBefore, msg.SenderBalanceAfter, &msg.Remarks)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("create sender transaction for handle transfer: %w", err)
	}

	// insert into table transaction for receiver
	recStatus := repo.FilterStatusSuccess
	recType := repo.FilterTransactionCredit
	recTx, err := t.txService.Create(tx, recStatus, msg.ReceiverId, recType, msg.Amount, msg.ReceiverBalancerBefore, msg.ReceiverBalanceAfter, &msg.Remarks)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("create sender transaction for handle transfer: %w", err)
	}

	// insert into table transfer
	_, err = t.create(tx, msg.Id, sendTx.Id, recTx.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("create transfer for handle transfer: %w", err)
	}

	// update balance sender
	_, err = t.userService.UpdateBalance(tx, msg.SenderId, msg.SenderBalanceAfter)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update sender balance for handle transfer: %w", err)
	}

	// update balance receiver
	_, err = t.userService.UpdateBalance(tx, msg.ReceiverId, msg.ReceiverBalanceAfter)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update receiver balance for handle transfer: %w", err)
	}

	tx.Commit()

	return nil
}

func (t *TransferService) create(tx *gorm.DB, id, sendTxId, recTxId uuid.UUID) (*models.Transfer, error) {
	transfer := &models.Transfer{
		Id:                    id,
		SenderTransactionId:   sendTxId,
		ReceiverTransactionId: recTxId,
	}

	save, err := t.transferRepo.Create(tx, transfer)
	if err != nil {
		return nil, err
	}

	return save, nil
}
