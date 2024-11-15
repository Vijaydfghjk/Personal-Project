package routes

import (
	controllers "GST_billing_api/Controllers"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Authorization token missing"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ConnectRouter(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/login", controllers.Login)
	protected := r.Group("/")
	protected.Use(AuthMiddleware())

	protected.POST("/products", controllers.AddProduct)
	protected.GET("/products", controllers.SearchProduct)
	protected.POST("/Bill", controllers.GenerateBill)
	protected.POST("/protected-endpoint", controllers.ProtectedRoute)

}
