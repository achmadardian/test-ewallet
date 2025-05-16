package main

import (
	"achmadardian/test-ewallet/routes"
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

	// init routes
	routes.InitRoutes(r)

	log.Println("server running at http://localhost:8080")
	r.Run(":8080")
}
