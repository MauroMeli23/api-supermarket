package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token de la variable de entorno
		expectedToken := os.Getenv("TOKEN")

		// Obtener el token del encabezado de la solicitud
		providedToken := c.GetHeader("TOKEN")

		if providedToken != expectedToken {
			c.JSON(401, gin.H{"error": "Acceso no autorizado"})
			c.Abort()
			return
		}

		// Continuar con la ejecucion normal si el token es valido
		// El next nos permite continuar con los demas handlers existentes
		c.Next()
	}
}
