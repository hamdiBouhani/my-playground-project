package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost:5432/playground?sslmode=enable")
	if err != nil {
		panic(err)
	}

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get the current working directory: %v\n", err)
		panic(err)
	}

	// Create the absolute path to the JSON file
	sourceURL := filepath.Join(currentDir, "migrations", "sql")
	sourceURL = "file:" + sourceURL

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(sourceURL, "playground", driver)
	if err != nil {
		fmt.Printf("NewWithDatabaseInstance: %v\n", err)
		panic(err)
	}
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
