package services

import (
	"achmadardian/test-ewallet/models"
	"achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/requests"
	"achmadardian/test-ewallet/utils/errs"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepo
}

func NewUserService(userRepo *repositories.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Create(req *requests.UserRequest) (*models.User, error) {
	isRegistered, err := u.userRepo.IsPhoneRegistered(req.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("check phone registerd: %w", err)
	}
	if isRegistered {
		return nil, errs.ErrPhoneAlreadyUsed
	}

	uuid := uuid.New()
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(req.Pin), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hashed pin: %w", err)
	}

	user := &models.User{
		Id:          uuid,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Pin:         string(hashedPin),
	}

	save, err := u.userRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("save user: %w", err)
	}

	return save, nil
}
