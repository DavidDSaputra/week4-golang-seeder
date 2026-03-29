package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/DavidDSaputra/week4-golang-seeder/models"
)

var DB *gorm.DB

func InitDB() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("Koneksi database gagal")
	}

	DB = database

	// Auto migrate models
	DB.AutoMigrate(&models.User{}, &models.Product{})
}
