package main

import (
	"amplassignment/database"
	"amplassignment/handlers"
	"amplassignment/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	// Public routes
	public := r.Group("/public")
	{
		public.GET("/tasks", handlers.GetTasks)
	}

	// Protected routes
	protected := r.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/:id", handlers.GetTask)
		protected.POST("/", handlers.CreateTask)
		protected.PUT("/:id", handlers.UpdateTask)
		protected.DELETE("/:id", handlers.DeleteTask)
	}

	r.Run(":8080")
}
