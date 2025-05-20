package repositories

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	DB db.Database
}

func NewUserRepo(DB db.Database) *UserRepo {
	return &UserRepo{DB: DB}
}

func (u *UserRepo) Create(user *models.User) (*models.User, error) {
	if err := u.DB.Write().Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) IsPhoneRegistered(phone string) (bool, error) {
	var user int64

	if err := u.DB.Read().Model(&models.User{}).Where("phone_number = ?", phone).Count(&user).Error; err != nil {
		return false, err
	}

	return user > 0, nil
}

func (u *UserRepo) GetById(id uuid.UUID) (*models.User, error) {
	var user models.User

	err := u.DB.Read().Select("id, first_name, last_name, phone_number, address, pin, balance").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetByPhone(phone string) (*models.User, error) {
	var user models.User

	err := u.DB.Read().
		Select("id, first_name, last_name, phone_number, address, pin, balance").
		Where("phone_number = ?", phone).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetBalance(tx *gorm.DB, id uuid.UUID) (*models.User, error) {
	var user models.User

	DB := tx
	if tx == nil {
		DB = u.DB.Read()
	}

	err := DB.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Select("balance").
		First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) UpdateBalance(tx *gorm.DB, id uuid.UUID, balance int64) (*models.User, error) {
	var user models.User

	DB := tx
	if tx == nil {
		DB = u.DB.Write()
	}

	err := DB.Model(&models.User{}).Where("id = ?", id).UpdateColumn("balance", balance).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) Update(user *models.User, id uuid.UUID) (*models.User, error) {
	if err := u.DB.Write().Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
