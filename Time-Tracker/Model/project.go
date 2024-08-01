package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"` 
	Tasks       []Task `json:"tasks"`
}
