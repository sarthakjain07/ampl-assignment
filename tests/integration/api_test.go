package integration

import (
	"amplassignment/database"
	"amplassignment/handlers"
	"amplassignment/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupIntegrationTest() *gin.Engine {
	database.InitDB()
	database.DB.AutoMigrate(&models.Task{})

	r := gin.Default()
	r.GET("/public/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTask)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	return r
}

func TestGetTasksIntegration(t *testing.T) {
	r := setupIntegrationTest()

	req, _ := http.NewRequest("GET", "/public/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTaskIntegration(t *testing.T) {
	r := setupIntegrationTest()

	task := models.Task{
		Title:       "Integration Test Task",
		Description: "Integration Test Description",
		Status:      "pending",
	}
	jsonValue, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("X-API-Key", "your-secret-api-key")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateTaskIntegration(t *testing.T) {
	r := setupIntegrationTest()

	// Create a task first
	task := models.Task{
		Title:       "Integration Test Task",
		Description: "Integration Test Description",
		Status:      "pending",
	}
	database.DB.Create(&task)

	// Update the task
	updatedTask := models.Task{
		Title:       "Updated Integration Task",
		Description: "Updated Integration Description",
		Status:      "in-progress",
	}
	jsonValue, _ := json.Marshal(updatedTask)

	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("X-API-Key", "your-secret-api-key")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTaskIntegration(t *testing.T) {
	r := setupIntegrationTest()

	// Create a task first
	task := models.Task{
		Title:       "Integration Test Task",
		Description: "Integration Test Description",
		Status:      "pending",
	}
	database.DB.Create(&task)

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	req.Header.Set("X-API-Key", "your-secret-api-key")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
