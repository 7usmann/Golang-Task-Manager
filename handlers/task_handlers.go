package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/7usmann/Golang-Task-Manager/db" // Import the db package
	"github.com/7usmann/Golang-Task-Manager/models"

	"github.com/gorilla/mux"
)

// Get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query(context.Background(), "SELECT id, title, description, completed FROM tasks")
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			http.Error(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	json.NewEncoder(w).Encode(tasks)
}

// Get a single task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	err := db.Conn.QueryRow(context.Background(), "SELECT id, title, description, completed FROM tasks WHERE id = $1", params["id"]).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// Create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Insert into PostgreSQL
	query := `INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3) RETURNING id`
	err = db.Conn.QueryRow(context.Background(), query, task.Title, task.Description, task.Completed).Scan(&task.ID)
	if err != nil {
		http.Error(w, "Failed to create task: "+err.Error(), http.StatusInternalServerError) // Log the error message
		return
	}

	json.NewEncoder(w).Encode(task)
}

// Update an existing task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get the task ID from the URL

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Prepare SQL query to update the task in the database
	query := `UPDATE tasks SET title=$1, description=$2, completed=$3 WHERE id=$4`
	_, err = db.Conn.Exec(context.Background(), query, task.Title, task.Description, task.Completed, params["id"])
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

// Delete a task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get the task ID from the URL

	// Prepare SQL query to delete the task from the database
	query := `DELETE FROM tasks WHERE id=$1`
	_, err := db.Conn.Exec(context.Background(), query, params["id"])
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}
