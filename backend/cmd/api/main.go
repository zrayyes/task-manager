package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Service struct for dependency injection
type Service struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	// Future dependencies like database connections can be added here
}

// NewService initializes a new Service
func NewService() *Service {
	return &Service{
		infoLog:  log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// HelloWorld is a basic handler function
func (s *Service) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	// Initialize service with dependencies
	service := NewService()

	// Define routes
	e.GET("/", service.HelloWorld)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
