package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

// Initialize the PostgreSQL connection
func InitDB() error {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		Conn, err = pgx.Connect(context.Background(), dbURL)
		if err == nil {
			log.Println("Connected to PostgreSQL database!")
			return nil
		}
		log.Printf("Unable to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second) // wait before retrying
	}
	return fmt.Errorf("could not connect to database after %d attempts: %v", maxRetries, err)
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
