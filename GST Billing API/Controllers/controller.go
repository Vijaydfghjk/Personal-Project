package controllers

import (
	config "GST_billing_api/Config"
	models "GST_billing_api/Models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

func Login(c *gin.Context) {

	var loginDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if loginDetails.Username == "admin" && loginDetails.Password == "admin@345" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": loginDetails.Username,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.SetCookie("jwat_token", tokenString, 3600, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{"Message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully"})
}

func SearchProduct(c *gin.Context) {

	product_code := c.DefaultQuery("product_code", "")
	product_name := c.DefaultQuery("product_name", "")

	var product models.Product

	if product_code != "" {

		if err := config.DB.Where("product_code=?", product_code).First(&product).Error; err != nil {

			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
	} else if product_name != "" {

		if err := config.DB.Where("product_name LIKE ?", "%"+product_name+"%").First(&product).Error; err != nil {

			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

	}
	c.JSON(http.StatusOK, product)
}

func GenerateBill(c *gin.Context) {

	var billRequest struct {
		ProductCode string `json:"product_code"`
		Quantity    int    `json:"quantity"`
	}

	err := c.ShouldBindJSON(&billRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var product models.Product
	if err := config.DB.Where("product_code = ?", billRequest.ProductCode).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	gstAmount := (product.ProductPrice * float64(billRequest.Quantity)) * (product.ProductGst / 100)
	totalAmount := (product.ProductPrice * float64(billRequest.Quantity)) + gstAmount

	bill := models.Bill{
		ProductCode: billRequest.ProductCode,
		Quantity:    billRequest.Quantity,
		TotalPrice:  product.ProductPrice * float64(billRequest.Quantity),
		GstAmount:   gstAmount,
		TotalAmount: totalAmount,
	}

	if err := config.DB.Create(&bill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't generate the bill"})
		return
	}

	c.JSON(http.StatusOK, bill)

}

func ProtectedRoute(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
		return
	}

	tokenString = tokenString[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Authorized"})
}
