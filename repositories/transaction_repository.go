package repositories

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"
	"achmadardian/test-ewallet/utils/pagination"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterColumn string
type FilterValue string

const (
	// filter by column
	FilterStatus          FilterColumn = "status"
	FilterTransactionType FilterColumn = "transaction_type"

	// filter by value
	// status
	FilterStatusPending FilterValue = "PENDING"
	FilterStatusSuccess FilterValue = "SUCCESS"
	FilterStatusFailed  FilterValue = "FAILED"

	// transaction_type
	FilterTransactionDebit  FilterValue = "DEBIT"
	FilterTransactionCredit FilterValue = "CREDIT"
)

type TransactionRepo struct {
	DB db.Database
}

func NewTransactionRepository(DB db.Database) *TransactionRepo {
	return &TransactionRepo{
		DB: DB,
	}
}

func (t *TransactionRepo) Create(tx *gorm.DB, tr *models.Transaction) (*models.Transaction, error) {
	DB := tx
	if tx == nil {
		DB = t.DB.Write()
	}

	if err := DB.Create(tr).Error; err != nil {
		return nil, err
	}

	return tr, nil
}

func (t *TransactionRepo) GetAllByUserId(tx *gorm.DB, userId uuid.UUID, p *pagination.Pagination) ([]models.Transaction, error) {
	var transaction []models.Transaction

	DB := tx
	if tx == nil {
		DB = t.DB.Read()
	}

	err := DB.
		Select("id, status, user_id, transaction_type, amount, remarks, balance_before, balance_after, created_at").
		Where("user_id = ?", userId).
		Where("status", FilterStatusSuccess).
		Preload("User").
		Scopes(pagination.Paginate(p)).
		Order("created_at DESC").
		Find(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, err
}
