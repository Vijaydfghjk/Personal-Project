package controller

import (
	"net/http"
	model "timetracker/Model"
	"timetracker/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB


func Register(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Login handles user login and JWT generation
func Login(c *gin.Context) {
	var input model.LoginInput
	var user model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
