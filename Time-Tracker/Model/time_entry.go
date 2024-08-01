package model

import (
	"gorm.io/gorm"
)

type TimeEntry struct {
	gorm.Model
	Duration int  `json:"duration" binding:"required"`
	TaskID   uint `json:"task_id"` 
	Task     Task `json:"task"`
}
