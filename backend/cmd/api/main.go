package main

import (
	"database/sql"
	"log"

	"github.com/zrayyes/task-manager/internal/repositories"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func main() {
	// Setup database connection
	connStr := "postgres://task_user:task_password@localhost:5432/task_manager"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Ping the database to check if it's alive
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	e := echo.New()

	// Initialize a new Service
	service := NewService()

	// Register routes
	service.RegisterRoutes(e)

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			service.infoLog.Printf("REQUEST: uri: %v, status: %v\n", v.URI, v.Status)
			return nil
		},
	}))

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
