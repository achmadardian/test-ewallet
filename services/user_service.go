package services

import (
	"achmadardian/test-ewallet/models"
	"achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/requests"
	"achmadardian/test-ewallet/responses"
	"achmadardian/test-ewallet/utils/errs"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo    *repositories.UserRepo
	authService *AuthService
}

func NewUserService(userRepo *repositories.UserRepo, authService *AuthService) *UserService {
	return &UserService{
		userRepo:    userRepo,
		authService: authService,
	}
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

func (u *UserService) Login(req *requests.UserLoginRequest) (*responses.LoginResponse, error) {
	user, err := u.userRepo.GetByPhone(req.PhoneNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrInvalidLogin
		}

		return nil, fmt.Errorf("get by phone: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(req.Pin)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, errs.ErrInvalidLogin
		}

		return nil, fmt.Errorf("compare pin: %w", err)
	}

	accToken, err := u.authService.GenerateToken(user.Id, TypeAccessToken, 15*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("access token: %w", err)
	}

	refToken, err := u.authService.GenerateToken(user.Id, TypeRefreshToken, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("refresh token: %w", err)
	}

	login := responses.LoginResponse{
		AccessToken:  accToken,
		RefreshToken: refToken,
		User: responses.UserLoginResponse{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}

	return &login, nil
}

func (u *UserService) UpdateUser(req *requests.UserUpdateRequest, id uuid.UUID) (*models.User, error) {
	user, err := u.userRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	if user == nil {
		return nil, errs.ErrDataNotFound
	}

	// set update
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Address = req.Address

	update, err := u.userRepo.Update(user, id)
	if err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	return update, nil
}

func (u *UserService) GetBalance(tx *gorm.DB, userId uuid.UUID) (*models.User, error) {
	balance, err := u.userRepo.GetBalance(tx, userId)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (u *UserService) GetAccount(userId uuid.UUID) (*models.User, error) {
	acc, err := u.userRepo.GetById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("get account: %w", err)
	}

	return acc, nil
}

func (u *UserService) UpdateBalance(tx *gorm.DB, userId uuid.UUID, balance int64) (*models.User, error) {
	update, err := u.userRepo.UpdateBalance(tx, userId, balance)
	if err != nil {
		return nil, err
	}

	return update, nil
}
