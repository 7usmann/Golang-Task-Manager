package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/7usmann/Golang-Task-Manager/db" // Import the db package
	"github.com/7usmann/Golang-Task-Manager/models"

	"github.com/gorilla/mux"
)

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	TaskType    string `json:"task_type"`
	TaskDate    string `json:"task_date"` // Expected in format YYYY-MM-DD
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Parse TaskDate to time.Time format
	taskDate, err := time.Parse("2006-01-02", req.TaskDate)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	// Insert the task into the database
	query := `INSERT INTO tasks (title, description, completed, task_date, task_type) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var taskID int
	err = db.Conn.QueryRow(context.Background(), query, req.Title, req.Description, req.Completed, taskDate, req.TaskType).Scan(&taskID)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	// Return a success response with the created task's ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      taskID,
		"message": "Task created successfully",
	})
}

// Get all tasks
func GetTasksByDate(w http.ResponseWriter, r *http.Request) {
	// Extract date parameters from the URL
	params := mux.Vars(r)
	day := params["day"]
	month := params["month"]
	year := params["year"]

	// Format the date to match the database format (e.g., YYYY-MM-DD)
	formattedDate := fmt.Sprintf("%s-%s-%s", year, month, day)

	// Query the database for tasks on the specified date
	rows, err := db.Conn.Query(context.Background(),
		"SELECT id, title, description, completed, task_date, task_type FROM tasks WHERE task_date = $1", formattedDate)
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Collect tasks in a slice
	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.TaskDate, &task.TaskType)
		if err != nil {
			http.Error(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	// Encode tasks as JSON and write to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Get a single task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	err := db.Conn.QueryRow(context.Background(), "SELECT id, title, description, completed, task_date, task_type FROM tasks WHERE id = $1", params["id"]).Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.TaskDate, &task.TaskType)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
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
	query := `UPDATE tasks SET title=$1, description=$2, completed=$3, task_date=$4, task_type=$5 WHERE id=$4`
	_, err = db.Conn.Exec(context.Background(), query, task.Title, task.Description, task.Completed, task.TaskDate, task.TaskType, params["id"])
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
