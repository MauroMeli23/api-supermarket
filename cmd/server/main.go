package main

import (
	"github.com/MauroMeli23/api-supermarket/cmd/server/handler"
	"github.com/MauroMeli23/api-supermarket/cmd/server/middleware"
	"github.com/MauroMeli23/api-supermarket/internal/domain"
	"github.com/MauroMeli23/api-supermarket/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var Products []domain.Product

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("La variable de entorno TOKEN no está definida")
	}

	// Cargar productos
	loadedProducts, err := product.LoadProducts()
	if err != nil {
		panic(err)
	}
	Products = loadedProducts
	// Inicializar el enrutador Gin
	r := gin.Default()
	r.Use(middleware.AuthMiddleware)

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
