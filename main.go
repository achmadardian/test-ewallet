package main

import (
	"achmadardian/test-ewallet/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// init gin
	r := gin.Default()

	// init routes
	routes.InitRoutes(r)

	log.Println("server running at http://localhost:8080")
	r.Run(":8080")
}
