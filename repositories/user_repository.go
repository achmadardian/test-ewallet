package repositories

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/models"
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
