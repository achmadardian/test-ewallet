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

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (a *AuthHandler) RefreshToken(c *gin.Context) {
	var req requests.RefreshToken

	if err := c.ShouldBindJSON(&req); err != nil {
		errVal := validate.ExtractValidationErrors(err)
		responses.BadRequest(c, errVal, errs.ErrValidationError.Error())
		return
	}

	renewToken, err := a.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidToken) {
			responses.Unauthorized(c)
			return
		}

		log.Printf("error create refresh token: %v", err)
		responses.InternalServerError(c)
		return
	}

	responses.Ok(c, renewToken)
}
