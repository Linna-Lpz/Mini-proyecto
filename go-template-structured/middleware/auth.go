package middleware

import (
	"go-template/helpers"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Autorización es requerida"})
			c.Abort()
			return
		}

		// Bearer <token>
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		claims,err := helpers.ValidateToken(authHeader)

		if err != nil {
			log.Println("Error al validar el token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set(("claims"), claims)
		c.Next()
	}

}