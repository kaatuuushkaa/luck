package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"task4/internal/db"
	"task4/internal/handlers"
	"task4/internal/taskService"
	"task4/internal/web/tasks"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	repo := taskService.NewTaskRepository(database)
	service := taskService.NewTaskService(repo)

	handler := handlers.NewTaskHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
