package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"task4/internal/db"
	"task4/internal/handlers"
	"task4/internal/taskService"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.POST("/tasks", taskHandlers.PostHandler)
	e.GET("/tasks", taskHandlers.GetHandler)
	e.PATCH("/tasks/:id", taskHandlers.PatchHandler)
	e.DELETE("/tasks/:id", taskHandlers.DeleteHandler)

	e.Start("localhost:8080")
}
