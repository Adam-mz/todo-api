package handlers

import (
	"net/http"
	"todo-api/terminal/models"
	"todo-api/terminal/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	createdUser, err := h.service.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, createdUser)
}
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get users")
	}
	return c.JSON(http.StatusOK, users)
}
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON")
	}

	updatedUser, err := h.service.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, updatedUser)
}
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.NoContent(http.StatusNoContent)
}
