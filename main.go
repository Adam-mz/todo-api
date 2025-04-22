package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to batabase: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

}

type Task struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Task    string `json:"task"`
	Is_Done bool   `json:"is_done"`
}

func taskGetHandler(c echo.Context) error {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}
	if len(tasks) == 0 {
		return c.JSON(http.StatusNotFound, "No tasks found")
	}
	return c.JSON(http.StatusOK, tasks)
}

func taskPostHandler(c echo.Context) error {
	var body Task
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	db.Create(&body)
	return c.JSON(http.StatusCreated, body)
}

func taskPutHandler(c echo.Context) error {
	id := c.Param("id")
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON")
	}

	db.Save(&task)
	return c.JSON(http.StatusOK, task)
}

func taskDeleteHandler(c echo.Context) error {
	id := c.Param("id")
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	db.Delete(&task)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/task", taskPostHandler)
	e.GET("/task", taskGetHandler)
	e.PUT("/task/:id", taskPutHandler)
	e.DELETE("/task/:id", taskDeleteHandler)
	e.Start(":8080")

}
