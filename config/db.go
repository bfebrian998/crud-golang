package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

func ConfigDatabase() *sql.DB {
	connStr := "user=postgres password=mysecretpassword dbname=crud_db host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Database connection successful!")
	return db
}

func InitDB() (*sql.DB, error) {
	db := ConfigDatabase()
	if db == nil {
		fmt.Println("error database")
	}
	return db, nil
}
