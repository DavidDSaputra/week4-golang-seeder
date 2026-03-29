package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/DavidDSaputra/week4-golang-seeder/config"
	"github.com/DavidDSaputra/week4-golang-seeder/models"
)

func main() {
	godotenv.Load()
	config.InitDB()

	products := []models.Product{
		{Name: "Nasi Goreng", Price: 20000, Stock: 10},
		{Name: "Mie Ayam", Price: 15000, Stock: 20},
		{Name: "Es Teh", Price: 5000, Stock: 50},
	}

	for _, product := range products {
		config.DB.Create(&product)
	}

	log.Println("Seeder berhasil dijalankan")
}
