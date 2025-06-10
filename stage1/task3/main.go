package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}
}

type Task struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST", http.StatusMethodNotAllowed)
		return
	}
	var req Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := db.Create(&req).Error; err != nil {
		http.Error(w, fmt.Sprintf("Could not create task: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET", http.StatusMethodNotAllowed)
		return
	}
	var tasks []Task

	if err := db.Find(&tasks).Error; err != nil {
		http.Error(w, fmt.Sprintf("Could not get tasks: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Only PATCH", http.StatusMethodNotAllowed)
		return
	}
	var req Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	var existing Task
	if err := db.First(&existing, "id = ?", req.ID).Error; err != nil {
		http.Error(w, fmt.Sprintf("Could not find expression: %v", err), http.StatusBadRequest)
		return
	}

	existing.Task = req.Task
	existing.IsDone = req.IsDone
	if err := db.Save(&existing).Error; err != nil {
		http.Error(w, fmt.Sprintf("Could not update task: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existing)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
	}

	if err := db.Delete(&Task{}, req.ID).Error; err != nil {
		http.Error(w, fmt.Sprintf("Could not delete task: %v", err), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()
	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/patch", PatchHandler)
	http.HandleFunc("/delete", DeleteHandler)
	http.ListenAndServe("localhost:8080", nil)
}
