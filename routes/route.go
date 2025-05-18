package routes

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/handlers"
	"achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, DB db.Database) {
	// routes
	api := r.Group("api")
	{
		// instantiate
		hc := handlers.NewHealthcheck()
		// user
		userRepo := repositories.NewUserRepo(DB)
		userService := services.NewUserService(userRepo)
		userHandler := handlers.NewUserHandler(userService)

		// healthcheck
		api.GET("/", hc.GetHealthcheck)

		// user
		api.POST("/register", userHandler.CreateUser)
		api.PUT("/update-profile/:id", userHandler.UpdateUser)
	}
}
