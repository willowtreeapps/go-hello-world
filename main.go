package main

import (
	"net/http"
	"strconv"
	"sync"

	_ "go-hello-world/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Task Manager API
// @version         1.0
// @description     A simple task manager API written in Go
// @host           localhost:8080
// @BasePath       /api/v1

// Task represents a simple todo task
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// TaskStore manages task storage with thread safety
type TaskStore struct {
	sync.RWMutex
	tasks []Task
}

// NewTaskStore creates a new TaskStore with initial tasks
func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: []Task{
			{ID: 1, Title: "Learn Go basics", Completed: false},
			{ID: 2, Title: "Write some code", Completed: false},
			{ID: 3, Title: "Build something cool", Completed: false},
		},
	}
}

var store = NewTaskStore()

// @Summary     Get all tasks
// @Description Get a list of all tasks
// @Tags        tasks
// @Produce     json
// @Success     200 {array}  Task
// @Router      /tasks [get]
func getTasks(c *gin.Context) {
	store.RLock()
	defer store.RUnlock()
	c.JSON(http.StatusOK, store.tasks)
}

// @Summary     Get a task by ID
// @Description Get details of a specific task
// @Tags        tasks
// @Produce     json
// @Param       id  path      int  true  "Task ID"
// @Success     200 {object}  Task
// @Failure     404 {object}  string
// @Router      /tasks/{id} [get]
func getTask(c *gin.Context) {
	store.RLock()
	defer store.RUnlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for _, task := range store.tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// @Summary     Mark a task as complete
// @Description Mark a specific task as completed
// @Tags        tasks
// @Produce     json
// @Param       id  path      int  true  "Task ID"
// @Success     200 {object}  Task
// @Failure     404 {object}  string
// @Router      /tasks/{id}/complete [post]
func completeTask(c *gin.Context) {
	store.Lock()
	defer store.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i := range store.tasks {
		if store.tasks[i].ID == id {
			store.tasks[i].Completed = true
			c.JSON(http.StatusOK, store.tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func main() {
	r := gin.Default()

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", getTasks)
			tasks.GET("/:id", getTask)
			tasks.POST("/:id/complete", completeTask)
		}
	}

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
