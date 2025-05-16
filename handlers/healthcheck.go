package handlers

import (
	"achmadardian/test-ewallet/responses"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct{}

func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}

func (h *Healthcheck) GetHealthcheck(c *gin.Context) {
	responses.Ok(c, "healthy")
}
