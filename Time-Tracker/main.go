package main

import (
	"log"
	controller "timetracker/Controllers"
	model "timetracker/Model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=postgres password=Vijay@123 dbname=myfile port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Project{}, &model.Task{}, &model.TimeEntry{})

	// Initialize the global DB variable
	controller.DB = db

	router := gin.Default()

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	router.Use(controller.AuthMiddleware())
	{
		router.POST("/projects", controller.CreateProject)
		router.PUT("/projects/:id", controller.UpdateProject)
		router.DELETE("/projects/:id", controller.DeleteProject)
		router.GET("/projects/:id", controller.GetProjectByID)
		router.GET("/projects", controller.GetAllProjects)
	}

	router.Run(":8080")
}
