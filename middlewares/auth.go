package middlewares

import (
	"achmadardian/test-ewallet/responses"
	"achmadardian/test-ewallet/services"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Auth(a *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		prefix := "Bearer "

		if !strings.HasPrefix(authorization, prefix) {
			responses.Unauthorized(c)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authorization, prefix)
		if token == " " {
			responses.Unauthorized(c)
			c.Abort()
			return
		}

		// validate access token
		claim, err := a.ValidateToken(token)
		if err != nil {
			responses.Unauthorized(c)
			log.Print("error: ", err)
			c.Abort()
			return
		}

		if claim.TokenType != services.TypeAccessToken {
			responses.Unauthorized(c)
			c.Abort()
			return
		}

		userId, err := uuid.Parse(claim.Subject)
		if err != nil {
			log.Printf("error parse uuid: %v", err)
			responses.InternalServerError(c)
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
