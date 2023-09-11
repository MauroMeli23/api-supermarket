package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func AuthMiddleware(c *gin.Context) {
	// Obtén el token de la variable de entorno
	expectedToken := os.Getenv("TOKEN")

	// Obtiene el token del encabezado de la solicitud
	providedToken := c.GetHeader("TOKEN")
	fmt.Println("TOKENNN", providedToken)
	if providedToken != expectedToken {
		c.JSON(401, gin.H{"error": "Acceso no autorizado"})
		c.Abort()
		return
	}

	// Continuar con la ejecución normal si el token es válido
	c.Next()
}
