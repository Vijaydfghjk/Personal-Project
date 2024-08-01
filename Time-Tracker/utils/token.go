package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("your_secret_key")

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
