package controller

import (
	"net/http"
	model "timetracker/Model"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var task model.Task
	id := c.Param("id")
	if err := DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&model.Task{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func GetTasks(c *gin.Context) {
	var tasks []model.Task
	projectID := c.Query("project_id")
	if err := DB.Where("project_id = ?", projectID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
