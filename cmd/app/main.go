package main

import (
	"log"
	"todo-api/terminal/database"
	"todo-api/terminal/handlers"
	"todo-api/terminal/repository"
	"todo-api/terminal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := database.InitDB()

	repo := repository.NewTaskRepository(db)
	srv := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(srv)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.POST("/tasks", handler.PostTasks)
	e.GET("/tasks", handler.GetAllTasks)
	e.GET("/tasks/:id", handler.GetTaskByID)
	e.PUT("/tasks/:id", handler.PatchTasksId)
	e.DELETE("/tasks/:id", handler.DeleteTask)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start: %v", err)
	}

}
