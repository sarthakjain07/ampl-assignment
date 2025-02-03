package handlers

import (
	"amplassignment/database"
	"amplassignment/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	database.InitDB()
	database.DB.AutoMigrate(&models.Task{})

	r := gin.Default()
	r.GET("/public/tasks", GetTasks)
	r.GET("/tasks/:id", GetTask)
	r.POST("/tasks", CreateTask)
	r.PUT("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)

	return r
}

func TestGetTasks(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/public/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTask(t *testing.T) {
	r := setupTestRouter()

	task := models.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}
	jsonValue, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("X-API-Key", "your-secret-api-key")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateTask(t *testing.T) {
	r := setupTestRouter()

	// Create a task first
	task := models.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}
	database.DB.Create(&task)

	// Update the task
	updatedTask := models.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "in-progress",
	}
	jsonValue, _ := json.Marshal(updatedTask)

	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("X-API-Key", "your-secret-api-key")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	r := setupTestRouter()

	// Create a task first
	task := models.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}
	database.DB.Create(&task)

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	req.Header.Set("X-API-Key", "your-secret-api-key")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
