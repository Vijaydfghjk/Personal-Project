package controller

import (
	"net/http"
	model "timetracker/Model"

	"github.com/gin-gonic/gin"
)

// Global DB variable
//var DB *gorm.DB

// CreateProject handles the creation of a new project
func CreateProject(c *gin.Context) {
	var project model.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	var user model.User
	if err := DB.First(&user, project.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Create the project
	if err := DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// UpdateProject handles updating an existing project
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var project model.Project
	if err := DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// DeleteProject handles deleting a project by ID
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&model.Project{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting project"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}

// GetProjectByID retrieves a project by its ID
func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	var project model.Project
	if err := DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// GetAllProjects retrieves all projects
func GetAllProjects(c *gin.Context) {
	var projects []model.Project
	if err := DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}
