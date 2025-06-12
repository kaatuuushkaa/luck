package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"task4/internal/db"
	"task4/internal/handlers"
	"task4/internal/taskService"
	"task4/internal/userService"
	"task4/internal/web/tasks"
	"task4/internal/web/users"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	repoTask := taskService.NewTaskRepository(database)
	serviceTask := taskService.NewTaskService(repoTask)

	repoUser := userService.NewUserRepository(database)
	serviceUser := userService.NewUserService(repoUser)

	handlerTask := handlers.NewTaskHandler(serviceTask)

	handlerUser := handlers.NewUserHandler(serviceUser)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandlerTask := tasks.NewStrictHandler(handlerTask, nil)
	tasks.RegisterHandlers(e, strictHandlerTask)

	strictHandlerUser := users.NewStrictHandler(handlerUser, nil)
	users.RegisterHandlers(e, strictHandlerUser)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
