package repositories

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"

	"gorm.io/gorm"
)

type TopupRepo struct {
	DB db.Database
}

func NewTopupRepository(DB db.Database) *TopupRepo {
	return &TopupRepo{
		DB: DB,
	}
}

func (t *TopupRepo) Create(tx *gorm.DB, topup *models.Topup) (*models.Topup, error) {
	DB := tx
	if tx == nil {
		DB = t.DB.Write()
	}

	if err := DB.Create(topup).Error; err != nil {
		return nil, err
	}

	return topup, nil
}
