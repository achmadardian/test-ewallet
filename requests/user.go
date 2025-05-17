package requests

import (
	"github.com/google/uuid"
)

type UserRequest struct {
	Id          uuid.UUID `json:"-"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Address     string    `json:"address" binding:"required"`
	Pin         string    `json:"pin" binding:"required"`
}
