package routes

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/handlers"
	"achmadardian/test-ewallet/middlewares"
	"achmadardian/test-ewallet/queue"
	"achmadardian/test-ewallet/queue/consumer"
	"achmadardian/test-ewallet/repositories"
	"achmadardian/test-ewallet/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, DB db.Database, rb *queue.RabbitMqClient) {
	// routes
	api := r.Group("api")
	{
		// instantiate
		hc := handlers.NewHealthcheck()
		// auth
		authService := services.NewAuthService()
		authHandler := handlers.NewAuthHandler(authService)

		// user
		userRepo := repositories.NewUserRepo(DB)
		userService := services.NewUserService(userRepo, authService)
		userHandler := handlers.NewUserHandler(userService)

		// transaction
		txRepo := repositories.NewTransactionRepository(DB)
		txService := services.NewTransactionService(txRepo)
		txHandler := handlers.NewTransactionHandler(txService)

		// topup
		topupRepo := repositories.NewTopupRepository(DB)
		topupService := services.NewTopService(topupRepo, userService, txService, DB)
		topupHandler := handlers.NewTopupHandler(topupService)

		// payments
		paymentRepo := repositories.NewPaymentRepository(DB)
		paymentService := services.NewPaymentService(paymentRepo, userService, txService, DB)
		paymentHandler := handlers.NewPaymentHandler(paymentService)

		// transfers
		transferRepo := repositories.NewTransferRepo(DB)
		transferService := services.NewTransferService(transferRepo, userService, rb, DB, txService)
		transferHandler := handlers.NewTransferHandler(transferService)

		// worker
		worker := consumer.NewTransferConsumer(rb, transferService)
		worker.ConsumeTransfer("transfer")

		// healthcheck
		api.GET("/", hc.GetHealthcheck)

		// user
		api.POST("/register", userHandler.CreateUser)
		api.POST("/login", userHandler.Login)
		api.POST("/refresh-token", authHandler.RefreshToken)

		// middleware
		api.Use(middlewares.Auth(authService))

		// protected routes
		api.PUT("/update-profile", userHandler.UpdateUser)
		api.GET("/account", userHandler.GetUserAccount)

		transaction := api.Group("/transactions")
		{
			transaction.GET("/", txHandler.GetAllTransactionByUserId)
		}

		// topup
		topup := api.Group("/topup")
		{
			topup.POST("/", topupHandler.Topup)
		}

		// payment
		payment := api.Group("/payments")
		{
			payment.POST("/", paymentHandler.Pay)
		}

		// transfer
		transfer := api.Group("/transfers")
		{
			transfer.POST("/", transferHandler.Transfer)
		}
	}
}
