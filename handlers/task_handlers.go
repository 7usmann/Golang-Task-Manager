package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/7usmann/Golang-Task-Manager/models"
	"github.com/gorilla/mux"
)

var tasks []models.Task

// Get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Get a single task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the task ID from the URL
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Task{})
}

// Create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)  // Skipping error check
	task.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe for production
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

// Update an existing task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			// Decode the new task data
			var updatedTask models.Task
			_ = json.NewDecoder(r.Body).Decode(&updatedTask)

			// Maintain the original task ID
			tasks[index].Title = updatedTask.Title
			tasks[index].Description = updatedTask.Description
			tasks[index].Completed = updatedTask.Completed

			json.NewEncoder(w).Encode(tasks[index])
			return
		}
	}
}

// Delete a task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)
}
