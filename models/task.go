package models

import (
	"context"
	"time"

	"github.com/7usmann/Golang-Task-Manager/db"
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

func GetTasksByDate(date string) ([]Task, error) {
	rows, err := db.Conn.Query(context.Background(), "SELECT id, title, description, completed, task_date, task_type FROM tasks WHERE task_date = $1", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.TaskDate, &task.TaskType); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
