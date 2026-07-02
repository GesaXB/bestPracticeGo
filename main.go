package main

import (
	"ginBestPractice/config"
	"ginBestPractice/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	server := gin.Default()
	routes.SetupRoutes(server, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := server.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
