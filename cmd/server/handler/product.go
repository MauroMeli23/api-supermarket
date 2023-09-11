package handler

import (
	"github.com/MauroMeli23/api-supermarket/internal/domain"
	"github.com/MauroMeli23/api-supermarket/internal/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllProducts(c *gin.Context, products []domain.Product) {
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context, products []domain.Product) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto no válido"})
		return
	}

	p, err := product.GetProductByID(id, products)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, p)
}

func GetProductByName(c *gin.Context, products []domain.Product) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Debe ingresar un nombre"})
		return
	}

	p, err := product.GetProductByName(name, products)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, p)
}

func AddNewProduct(c *gin.Context, products *[]domain.Product) error {
	var newProduct domain.Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	err := product.AddNewProduct(newProduct, products)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	return nil
}
