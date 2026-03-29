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

func CreateProduct(c *gin.Context) {
	var product models.Product
	c.ShouldBindJSON(&product)

	config.DB.Create(&product)
	c.JSON(201, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	config.DB.First(&product, id)

	c.ShouldBindJSON(&product)
	config.DB.Save(&product)

	c.JSON(200, product)
}
