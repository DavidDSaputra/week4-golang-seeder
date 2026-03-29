package handlers

import (
	"net/http"

	"github.com/DavidDSaputra/week4-golang-seeder/config"
	"github.com/DavidDSaputra/week4-golang-seeder/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}
