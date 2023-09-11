package main

import (
	"github.com/MauroMeli23/api-supermarket/cmd/server/handler"
	"github.com/MauroMeli23/api-supermarket/internal/domain"
	"github.com/MauroMeli23/api-supermarket/internal/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Products []domain.Product

func main() {
	// Cargar productos
	loadedProducts, err := product.LoadProducts()
	if err != nil {
		panic(err)
	}
	Products = loadedProducts
	// Inicializar el enrutador Gin
	r := gin.Default()

	// Definir las rutas
	r.GET("/products", func(c *gin.Context) {
		handler.GetAllProducts(c, Products)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		handler.GetProductByID(c, Products)
	})

	r.GET("/products/search", func(c *gin.Context) {
		handler.GetProductByName(c, Products)
	})

	r.POST("/products", func(c *gin.Context) {
		err := handler.AddNewProduct(c, &Products)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Producto añadido correctamente"})
	})

	// Iniciar el servidor
	r.Run(":8080")
}
