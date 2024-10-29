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

	// Run migrations
	db.RunMigrations()
	r := mux.NewRouter()

	// Route handlers
	// Route handlers
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/", handlers.HomePageHandler).Methods("GET")
	// Serve the HTML layout
	r.HandleFunc("/date/{day}/{month}/{year}", handlers.DatePageHandler).Methods("GET")
	// Fetch task data for the specified date via API
	r.HandleFunc("/api/tasks/{day}/{month}/{year}", handlers.GetTasksByDate).Methods("GET")

	// r.HandleFunc("/date/{date}", handlers.CreateTaskByDate).Methods("POST")        // To add a task for a specific date
	// r.HandleFunc("/date/{date}/{id}", handlers.UpdateTaskByDate).Methods("PUT")    // To update a specific task by ID on a date
	// r.HandleFunc("/date/{date}/{id}", handlers.DeleteTaskByDate).Methods("DELETE") // To delete a specific task by ID on a date

	log.Println("Server starting on port 8080...")
	erro := http.ListenAndServe(":8080", r)
	if erro != nil {
		log.Fatal("ListenAndServe: ", erro)
	}
}
