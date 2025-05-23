package handlers

import (
	"net/http"
	"strconv"
	"todo-api/terminal/models"
	"todo-api/terminal/service"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(s service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) PostTasks(c echo.Context) error {
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	createdTask, err := h.service.PostTasks(&task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create task")
	}

	return c.JSON(http.StatusCreated, createdTask)
}
func (h *TaskHandler) GetTasksByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}
	tasks, err := h.service.GetTasksByUserID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not get tasks")
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get tasks")
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	id := c.Param("id")
	task, err := h.service.GetTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) PatchTasksId(c echo.Context) error {
	id := c.Param("id")
	task, err := h.service.GetTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON")
	}

	updatedTask, err := h.service.PatchTasksId(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update task")
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete task")
	}
	return c.NoContent(http.StatusNoContent)
}
