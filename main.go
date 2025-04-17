package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type requestBody struct {
	Task string `json:"task"`
}

var task string

func taskGetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, "+task)
}

func taskPostHandler(c echo.Context) error {
	var body requestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	task = body.Task
	return c.JSON(http.StatusOK, "Task created")
}

func main() {
	e := echo.New()

	e.POST("/task", taskPostHandler)
	e.GET("/task", taskGetHandler)

	e.Start(":8080")

}
