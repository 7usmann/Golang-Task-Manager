package main

import (
	"log"
	"net/http"

	"github.com/7usmann/Golang-Task-Manager/db"
	"github.com/7usmann/Golang-Task-Manager/handlers"

	"github.com/gorilla/mux"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	log.Println("This is a test log message")

	// Run migrations
	db.RunMigrations()
	r := mux.NewRouter()

	// Route handlers
	// Route handlers
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/api/tasks/month/{year}/{month}", handlers.GetTasksByMonth).Methods("GET")
	r.HandleFunc("/", handlers.HomePageHandler).Methods("GET")
	// Serve the HTML layout
	r.HandleFunc("/date/{day}/{month}/{year}", handlers.DatePageHandler).Methods("GET")
	// Fetch task data for the specified date via API
	r.HandleFunc("/api/tasks/{day}/{month}/{year}", handlers.GetTasksByDate).Methods("GET")
	r.HandleFunc("/api/tasks/{day}/{month}/{year}", handlers.CreateTask).Methods("POST") // To add a task for a specific date
	r.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/api/tasks/week", handlers.GetTasksByWeek).Methods("GET")
	r.HandleFunc("/api/tasks/day/{year}/{month}/{day}", handlers.GetTasksByDay).Methods("GET")

	log.Println("Server starting on port 8080...")
	erro := http.ListenAndServe(":8080", r)
	if erro != nil {
		log.Fatal("ListenAndServe: ", erro)
	}
}
