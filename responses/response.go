package responses

import (
	"achmadardian/test-ewallet/utils/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code       int                    `json:"code"`
	Status     string                 `json:"status"`
	Message    string                 `json:"message"`
	Data       any                    `json:"result,omitempty"`
	Errors     any                    `json:"errors,omitempty"`
	Pagination *pagination.Pagination `json:"pagination,omitempty"`
}

var (
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
)

var (
	MsgSuccess             = "ok"
	MsgCreated             = "created"
	MsgUpdated             = "updated"
	MsgDeleted             = "deleted"
	MsgBadRequest          = "bad request"
	MsgNotFound            = "not found"
	MsgUnauthorized        = "unauthorized"
	MsgInternalServerError = "internal server error"
)

func Ok(c *gin.Context, data any, message ...string) {
	msg := MsgSuccess
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Status:  StatusSuccess,
		Message: msg,
		Data:    data,
	})
}

func OkPage(c *gin.Context, data any, p *pagination.Pagination, message ...string) {
	msg := MsgSuccess
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Status:  StatusSuccess,
		Message: msg,
		Data:    data,
		Pagination: &pagination.Pagination{
			Page:     p.Page,
			PageSize: p.PageSize,
		},
	})
}

func Created(c *gin.Context, data any, message ...string) {
	msg := MsgCreated
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusCreated, ApiResponse{
		Code:    http.StatusCreated,
		Status:  StatusSuccess,
		Message: msg,
		Data:    data,
	})
}

func Updated(c *gin.Context, data any, message ...string) {
	msg := MsgUpdated
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Status:  StatusSuccess,
		Message: msg,
		Data:    data,
	})
}

func Deleted(c *gin.Context, message ...string) {
	msg := MsgDeleted
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Status:  StatusSuccess,
		Message: msg,
	})
}

func NotFound(c *gin.Context, message ...string) {
	msg := MsgNotFound
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusNotFound, ApiResponse{
		Code:    http.StatusNotFound,
		Status:  StatusFailed,
		Message: msg,
	})
}

func BadRequest(c *gin.Context, errors any, message ...string) {
	msg := MsgBadRequest
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusBadRequest, ApiResponse{
		Code:    http.StatusBadRequest,
		Status:  StatusFailed,
		Message: msg,
		Errors:  errors,
	})
}

func Unauthorized(c *gin.Context, message ...string) {
	msg := MsgUnauthorized
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusUnauthorized, ApiResponse{
		Code:    http.StatusUnauthorized,
		Status:  StatusFailed,
		Message: msg,
	})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, ApiResponse{
		Code:    http.StatusInternalServerError,
		Status:  StatusFailed,
		Message: MsgInternalServerError,
	})
}
