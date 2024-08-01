package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description"`
	ProjectID   uint        `json:"project_id"` 
	Project     Project     `json:"project"`
	TimeEntries []TimeEntry `json:"time_entries"`
}
