package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zrayyes/task-manager/internal/repositories"
)

type TaskHandler struct {
	taskRepository repositories.TaskRepositoryInterface
}

func NewTaskHandler(tr repositories.TaskRepositoryInterface) *TaskHandler {
	return &TaskHandler{
		taskRepository: tr,
	}
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	return c.String(http.StatusOK, "Get Tasks placeholder")
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	return c.String(http.StatusCreated, "Create Task placeholder")
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	return c.String(http.StatusOK, "Get Task placeholder")
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	return c.String(http.StatusOK, "Update Task placeholder")
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	return c.String(http.StatusOK, "Delete Task placeholder")
}
