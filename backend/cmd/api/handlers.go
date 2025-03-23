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
	tasks := h.taskRepository.GetAllTasks()
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	task := new(repositories.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := h.taskRepository.CreateTask(task); err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	id := c.Param("id")
	task, err := h.taskRepository.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	updatedTask := new(repositories.Task)
	if err := c.Bind(updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := h.taskRepository.UpdateTask(id, updatedTask); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if err := h.taskRepository.DeleteTask(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
