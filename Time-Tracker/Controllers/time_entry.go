package controller

import (
	"net/http"
	model "timetracker/Model"

	"github.com/gin-gonic/gin"
)

func CreateTimeEntry(c *gin.Context) {
	var timeEntry model.TimeEntry
	if err := c.ShouldBindJSON(&timeEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&timeEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating time entry"})
		return
	}

	c.JSON(http.StatusOK, timeEntry)
}

func UpdateTimeEntry(c *gin.Context) {
	var timeEntry model.TimeEntry
	id := c.Param("id")
	if err := DB.First(&timeEntry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Time entry not found"})
		return
	}

	if err := c.ShouldBindJSON(&timeEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Save(&timeEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating time entry"})
		return
	}

	c.JSON(http.StatusOK, timeEntry)
}

func DeleteTimeEntry(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&model.TimeEntry{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting time entry"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Time entry deleted"})
}

func GetTimeEntries(c *gin.Context) {
	var timeEntries []model.TimeEntry
	taskID := c.Query("task_id")
	if err := DB.Where("task_id = ?", taskID).Find(&timeEntries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching time entries"})
		return
	}
	c.JSON(http.StatusOK, timeEntries)
}
