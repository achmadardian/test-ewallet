package routes

import (
	"achmadardian/test-ewallet/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// routes
	api := r.Group("api")
	{
		// instantiate
		hc := handlers.NewHealthcheck()

		// healthcheck
		api.GET("/", hc.GetHealthcheck)
	}
}
