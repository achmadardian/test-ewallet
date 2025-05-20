package repositories

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	DB db.Database
}

func NewPaymentRepository(DB db.Database) *PaymentRepo {
	return &PaymentRepo{
		DB: DB,
	}
}

func (p *PaymentRepo) Create(tx *gorm.DB, payment *models.Payment) (*models.Payment, error) {
	DB := tx
	if tx == nil {
		DB = p.DB.Write()
	}

	if err := DB.Create(payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}
