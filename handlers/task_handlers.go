package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	params := mux.Vars(r)
	day := params["day"]
	month := params["month"]
	year := params["year"]

	// Format date as YYYY-MM-DD
	taskDate := fmt.Sprintf("%s-%s-%s", year, month, day)

	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert the task into the database
	query := `INSERT INTO tasks (title, description, completed, task_date, task_type) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var taskID int
	err := db.Conn.QueryRow(context.Background(), query, req.Title, req.Description, req.Completed, taskDate, req.TaskType).Scan(&taskID)
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID := params["id"]

	query := "DELETE FROM tasks WHERE id = $1"
	_, err := db.Conn.Exec(context.Background(), query, taskID)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TaskType    string `json:"task_type"`
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID := params["id"]

	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE tasks SET title = $1, description = $2, task_type = $3 WHERE id = $4"
	_, err := db.Conn.Exec(context.Background(), query, req.Title, req.Description, req.TaskType, taskID)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

// GetTasksByMonth retrieves all tasks for a specific month
func GetTasksByMonth(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTasksByMonth called")
	params := mux.Vars(r)
	year := params["year"]
	month := params["month"]

	// Log year and month values for debugging
	log.Printf("Fetching tasks for year: %s, month: %s\n", year, month)

	// Query tasks for the specified month and year
	query := `
        SELECT id, title, description, completed, task_date, task_type 
        FROM tasks 
        WHERE EXTRACT(YEAR FROM task_date) = $1 AND EXTRACT(MONTH FROM task_date) = $2
    `

	// Execute the query
	rows, err := db.Conn.Query(context.Background(), query, year, month)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process the rows
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.TaskDate, &task.TaskType); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			http.Error(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	// Log the number of tasks fetched for verification
	log.Printf("Number of tasks fetched: %d\n", len(tasks))

	// Encode tasks as JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Printf("Error encoding tasks to JSON: %v\n", err)
		http.Error(w, "Failed to encode tasks to JSON", http.StatusInternalServerError)
	}
}

// GetTasksByWeek retrieves all tasks for a specific week
// GetTasksByWeek retrieves all tasks within a specific week range
func GetTasksByWeek(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start")
	endDate := r.URL.Query().Get("end")

	query := `
        SELECT id, title, description, completed, task_date, task_type 
        FROM tasks 
        WHERE task_date BETWEEN $1 AND $2
    `

	rows, err := db.Conn.Query(context.Background(), query, startDate, endDate)
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.TaskDate, &task.TaskType); err != nil {
			http.Error(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTasksByDay retrieves all tasks for a specific day
func GetTasksByDay(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	year := params["year"]
	month := params["month"]
	day := params["day"]

	formattedDate := fmt.Sprintf("%s-%s-%s", year, month, day)

	rows, err := db.Conn.Query(context.Background(),
		"SELECT id, title, description, completed, task_date, task_type FROM tasks WHERE task_date = $1", formattedDate)
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.TaskDate, &task.TaskType)
		if err != nil {
			http.Error(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
