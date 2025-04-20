package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type requestBody struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var task []requestBody
var TaskId int = 1

func taskGetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &task)
}

func taskPostHandler(c echo.Context) error {
	var body requestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	task = append(task, requestBody{ID: strconv.Itoa(TaskId), Task: body.Task})
	TaskId++
	return c.JSON(http.StatusOK, "Task created")
}

func taskPutHandler(c echo.Context) error {
	var body requestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	var id = c.Param("id")
	for i, t := range task {
		if t.ID == id {
			task[i].Task = body.Task
			return c.JSON(http.StatusOK, "Task updated")
		}

	}
	return c.JSON(http.StatusNotFound, "Task not found")
}

func taskDeleteHandler(c echo.Context) error {
	var id = c.Param("id")
	for i, t := range task {
		if t.ID == id {
			task = append(task[:i], task[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Task not found")
}

func main() {
	e := echo.New()

	e.POST("/task", taskPostHandler)
	e.GET("/task", taskGetHandler)
	e.PUT("/task/:id", taskPutHandler)
	e.DELETE("/task/:id", taskDeleteHandler)
	e.Start(":8080")

}
