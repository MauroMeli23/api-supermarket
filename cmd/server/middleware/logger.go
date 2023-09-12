package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Llama a los siguientes middlewares
		c.Next()

		// Obtiene informacion sobre la solicitud
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		latency := time.Since(start)
		clientIP := c.ClientIP()

		// Registra la informacion de la solicitud
		fmt.Printf("[Request] %v | %3d | %12v | %s | %-7s %s\n",
			start.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
