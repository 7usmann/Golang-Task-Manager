package main

import (
	"log"
	"net/http"

	"github.com/7usmann/Golang-Task-Manager/db"
	"github.com/7usmann/Golang-Task-Manager/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db.InitDB()
	defer db.CloseDB()

	r := mux.NewRouter()

	// Route handlers
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/", handlers.HomePageHandler).Methods("GET")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
