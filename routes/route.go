package routes

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/handlers"
	"achmadardian/test-ewallet/middlewares"
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
		// auth
		authService := services.NewAuthService()

		// user
		userRepo := repositories.NewUserRepo(DB)
		userService := services.NewUserService(userRepo, authService)
		userHandler := handlers.NewUserHandler(userService)

		// healthcheck
		api.GET("/", hc.GetHealthcheck)

		// user
		api.POST("/register", userHandler.CreateUser)
		api.POST("/login", userHandler.Login)

		// middleware
		api.Use(middlewares.Auth(authService))

		// protected routes
		api.PUT("/update-profile", userHandler.UpdateUser)
	}
}
