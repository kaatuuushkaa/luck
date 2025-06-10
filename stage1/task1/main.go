package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string

type TaskRequest struct {
	Task string `json:"task"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req TaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
			return
		}
		task = req.Task
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

func main() {
	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/get", GetHandler)
	http.ListenAndServe("localhost:8080", nil)
}
