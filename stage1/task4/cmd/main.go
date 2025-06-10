package main

import (
	"log"
	"net/http"
	"task4/internal/db"
	"task4/internal/handlers"
	"task4/internal/taskService"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	http.HandleFunc("/post", taskHandlers.PostHandler)
	http.HandleFunc("/get", taskHandlers.GetHandler)
	http.HandleFunc("/patch", taskHandlers.PatchHandler)
	http.HandleFunc("/delete", taskHandlers.DeleteHandler)
	http.ListenAndServe("localhost:8080", nil)
}
