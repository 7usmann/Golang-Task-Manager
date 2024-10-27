package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

// Initialize the PostgreSQL connection
func InitDB() {
	var err error
	Conn, err = pgx.Connect(context.Background(), "postgres://usman:7usmann@localhost:5432/task_manager")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	log.Println("Connected to PostgreSQL database!")
}

// Close the PostgreSQL connection
func CloseDB() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}
