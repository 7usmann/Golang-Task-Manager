package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

// Initialize the PostgreSQL connection
func InitDB() error {
	var err error
	Conn, err = pgx.Connect(context.Background(), "postgres://postgres:7usmann@localhost:5432/task_manager")
	if err != nil {
		log.Println("Unable to connect to database:", err)
		return err
	}
	log.Println("Connected to PostgreSQL database!")
	return nil
}

// Close the PostgreSQL connection
func CloseDB() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}

func RunMigrations() {
	// Read SQL file content
	sqlFile, err := os.Open("migrations/create_tables.sql")
	if err != nil {
		log.Fatalf("Error opening SQL file: %v", err)
	}
	defer sqlFile.Close()

	sqlContent, err := ioutil.ReadAll(sqlFile)
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}

	// Execute SQL commands
	_, err = Conn.Exec(context.Background(), string(sqlContent))
	if err != nil {
		log.Fatalf("Error executing migration: %v", err)
	}
	fmt.Println("Database migration successful!")
}
