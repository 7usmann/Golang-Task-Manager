package models

import (
	"time"
)

// Task represents a task in the task manager
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	TaskDate    time.Time `json:"task_date"` // Make sure to import "time"
	TaskType    string    `json:"task_type"`
}
