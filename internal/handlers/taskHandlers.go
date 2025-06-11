package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"task4/internal/taskService"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) PostHandler(c echo.Context) error {
	var req taskService.Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := h.service.CreateTask(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, req)
}

func (h *TaskHandler) GetHandler(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PatchHandler(c echo.Context) error {
	id := c.Param("id")

	var req taskService.TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	task, err := h.service.UpdateTask(id, req.Task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
