package requests

import (
	"github.com/google/uuid"
)

type UserRequest struct {
	Id          uuid.UUID `json:"-"`
	FirstName   string    `json:"first_name" binding:"required,lte=50"`
	LastName    string    `json:"last_name" binding:"required,lte=50"`
	PhoneNumber string    `json:"phone_number" binding:"required,lte=20"`
	Address     string    `json:"address" binding:"required"`
	Pin         string    `json:"pin" binding:"required,lte=6"`
}

type UserUpdateRequest struct {
	Id        uuid.UUID `json:"-"`
	FirstName string    `json:"first_name" binding:"required,lte=50"`
	LastName  string    `json:"last_name" binding:"required,lte=50"`
	Address   string    `json:"address,omitempty"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,lte=20"`
	Pin         string `json:"pin" binding:"required,lte=6"`
}
