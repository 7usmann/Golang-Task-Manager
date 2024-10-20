package main

import (
	"log"
	"net/http"

	"github.com/7usmann/Golang-Task-Manager/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route Handlers
	r.HandleFunc("/", handlers.HomePageHandler).Methods("GET")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Start server
	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
