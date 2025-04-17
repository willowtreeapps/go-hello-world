package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new router
	r := gin.Default()

	// Setup routes
	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", getTasks)
			tasks.GET("/:id", getTask)
			tasks.POST("/:id/complete", completeTask)
		}
	}

	return r
}

func TestGetTasks(t *testing.T) {
	// Reset the store to a known state
	store = NewTaskStore()

	router := setupTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 3) // We expect 3 initial tasks
	assert.Equal(t, "Learn Go basics", response[0].Title)
}

func TestGetTask(t *testing.T) {
	// Reset the store to a known state
	store = NewTaskStore()

	router := setupTestRouter()

	t.Run("Valid Task ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/tasks/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, "Learn Go basics", response.Title)
	})

	t.Run("Invalid Task ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/tasks/999", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Non-numeric Task ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/tasks/abc", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCompleteTask(t *testing.T) {
	// Reset the store to a known state
	store = NewTaskStore()

	router := setupTestRouter()

	t.Run("Valid Task ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/tasks/1/complete", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 1, response.ID)
		assert.True(t, response.Completed)
	})

	t.Run("Invalid Task ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/tasks/999/complete", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Non-numeric Task ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/tasks/abc/complete", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
