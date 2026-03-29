package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/DavidDSaputra/week4-golang-seeder/config"
	"github.com/DavidDSaputra/week4-golang-seeder/routes"
)

func main() {
	// Load environment config
	godotenv.Load()

	// Initialize database
	config.InitDB()

	// Setup routes
	server := routes.SetupRouter()

	// Get port dari env atau default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("running on", port)
	server.Run(":" + port)
}
