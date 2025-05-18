package responses

import (
	"github.com/google/uuid"
)

type UserResponse struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Created_at  string    `json:"created_at"`
}
