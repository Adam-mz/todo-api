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

	tasksRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(tasksRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	usersRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(usersRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.POST("/tasks", taskHandler.PostTasks)
	e.GET("/tasks", taskHandler.GetAllTasks)
	e.GET("/tasks/:id", taskHandler.GetTaskByID)
	e.PUT("/tasks/:id", taskHandler.PatchTasksId)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)

	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
	e.GET("/users/:user_id/tasks", taskHandler.GetTasksByUserID)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start: %v", err)
	}

}
