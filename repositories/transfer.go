package repositories

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"

	"gorm.io/gorm"
)

type TransferRepo struct {
	DB db.Database
}

func NewTransferRepo(DB db.Database) *TransferRepo {
	return &TransferRepo{
		DB: DB,
	}
}

func (t *TransferRepo) Create(tx *gorm.DB, transfer *models.Transfer) (*models.Transfer, error) {
	DB := tx
	if tx == nil {
		DB = t.DB.Write()
	}

	if err := DB.Create(transfer).Error; err != nil {
		return nil, err
	}

	return transfer, nil
}
