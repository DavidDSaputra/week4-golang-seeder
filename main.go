package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/DavidDSaputra/week4-golang-seeder/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment config
	godotenv.Load()

	// Initialize database
	config.InitDB()

	// Create router
	server := gin.Default()

	// Test endpoint
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	// Get port dari env atau default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("running on", port)
	server.Run(":" + port)
}
