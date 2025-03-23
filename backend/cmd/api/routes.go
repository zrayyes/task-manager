package main

import (
	"github.com/labstack/echo/v4"
)

func (s *Service) RegisterRoutes(e *echo.Echo) {

	// Task CRUD routes
	e.GET("/tasks", s.taskHandler.GetTasks)
	e.POST("/tasks", s.taskHandler.CreateTask)
	e.GET("/tasks/:id", s.taskHandler.GetTask)
	e.PUT("/tasks/:id", s.taskHandler.UpdateTask)
	e.DELETE("/tasks/:id", s.taskHandler.DeleteTask)
}
