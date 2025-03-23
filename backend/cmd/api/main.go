package main

import (
	"log"
	"net/http"

	"github.com/zrayyes/task-manager/internal/repositories"

	"github.com/labstack/echo/v4"
)

// Service struct for dependency injection
type Service struct {
	infoLog     *log.Logger
	errorLog    *log.Logger
	taskHandler *TaskHandler
}

// NewService initializes a new Service
func NewService() *Service {
	tr := repositories.NewTaskRepository()
	return &Service{
		infoLog:     log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog:    log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		taskHandler: NewTaskHandler(tr),
	}
}

// HelloWorld is a basic handler function
func (s *Service) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	// Initialize a new Service
	service := NewService()

	// Register routes
	service.RegisterRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
