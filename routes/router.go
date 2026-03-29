package routes

import (
	"github.com/DavidDSaputra/week4-golang-seeder/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/products", handlers.GetProducts)
	r.POST("/products", handlers.CreateProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	return r
}
