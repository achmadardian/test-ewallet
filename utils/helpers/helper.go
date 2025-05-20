package helpers

import (
	"achmadardian/test-ewallet/utils/pagination"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserId(c *gin.Context) (uuid.UUID, bool) {
	val, ok := c.Get("user_id")
	if !ok {
		return uuid.Nil, false
	}

	userId, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, false
	}

	return userId, true
}

func GetQueryParamPagination(c *gin.Context) (*pagination.Pagination, error) {
	var page pagination.Pagination

	err := c.BindQuery(&page)
	if err != nil {
		return nil, err
	}

	return &page, nil
}
