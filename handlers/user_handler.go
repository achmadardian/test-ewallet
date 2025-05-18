package handlers

import (
	"achmadardian/test-ewallet/requests"
	"achmadardian/test-ewallet/responses"
	"achmadardian/test-ewallet/services"
	"achmadardian/test-ewallet/utils/errs"
	"achmadardian/test-ewallet/utils/validate"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var req requests.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errRes := validate.ExtractValidationErrors(err)
		responses.BadRequest(c, errRes, errs.ErrValidationError.Error())
		return
	}

	user, err := u.userService.Create(&req)
	if err != nil {
		if errors.Is(err, errs.ErrPhoneAlreadyUsed) {
			responses.BadRequest(c, gin.H{"phone_number": errs.ErrPhoneAlreadyUsed.Error()}, "validation error")
			return
		}

		log.Printf("failed to create user: %v", err)
		responses.InternalServerError(c)
		return
	}

	res := responses.UserResponse{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.PhoneNumber,
		Created_at:  user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	responses.Created(c, res)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var req requests.UserUpdateRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		errRes := validate.ExtractValidationErrors(err)
		responses.BadRequest(c, errRes, errs.ErrValidationError.Error())
		return
	}

	update, err := u.userService.UpdateUser(&req, id)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			responses.NotFound(c, errs.ErrDataNotFound.Error())
			return
		}

		log.Printf("failed to update user: %v", err)
		responses.InternalServerError(c)
		return
	}

	res := responses.UserUpdateRespons{
		Id:        update.Id,
		FirstName: update.FirstName,
		LastName:  update.LastName,
		Address:   update.Address,
		UpdatedAt: update.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	responses.Updated(c, res)
}
