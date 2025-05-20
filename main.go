package main

import (
	"achmadardian/test-ewallet/db"
	"achmadardian/test-ewallet/queue"
	"achmadardian/test-ewallet/routes"
	"achmadardian/test-ewallet/utils/validate"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// get .env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, relying on OS variables")
	}

	// init gin
	r := gin.Default()

	// init DB
	DB := db.InitDB()

	// init rabbit
	rabbitmq := queue.Init()
	defer rabbitmq.Conn.Close()
	defer rabbitmq.Chann.Close()

	// init routes
	routes.InitRoutes(r, DB, rabbitmq)

	// init translator
	validate.InitTranslator()

	log.Println("server running at http://localhost:8080")
	r.Run(":8080")
}
