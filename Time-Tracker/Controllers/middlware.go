package controller

import (
	"net/http"
	"strings"
	"timetracker/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware checks for JWT in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenParts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.JwtKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims["user_id"])
		c.Next()
	}
}
