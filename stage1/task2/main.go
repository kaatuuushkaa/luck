package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task []string

type TaskRequest struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req TaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
			return
		}
		task = append(task, req.Task)
	} else {
		fmt.Fprintln(w, "Only POST")
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "hello, %s", task)
	} else {
		fmt.Fprintln(w, "Only GET")
	}
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		var req TaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		}
		if req.ID <= 0 || req.ID > len(task) {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		task[req.ID-1] = req.Task
	} else {
		fmt.Fprintln(w, "Only PATCH")
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		var req struct {
			ID int `json:"id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		}
		if req.ID <= 0 || req.ID > len(task) {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		task = append(task[:req.ID-1], task[req.ID:]...)
	} else {
		fmt.Fprintln(w, "Only DELETE")
	}
}

func main() {
	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/patch", PatchHandler)
	http.HandleFunc("/delete", DeleteHandler)
	http.ListenAndServe("localhost:8080", nil)
}
