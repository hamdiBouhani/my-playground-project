package config

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

type MigrateCfg struct {
}

func NewMigrateCfg() (*MigrateCfg, error) {
	return &MigrateCfg{}, nil
}

func (mg *MigrateCfg) Migrate() error {
	dns := "postgres://postgres:postgres@localhost:5432/playground?sslmode=disable"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return err
	}
	defer db.Close()

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get the current working directory: %v\n", err)
		return err
	}

	sourceURL := filepath.Join(currentDir, "migrations")
	sourceURL = "file:" + sourceURL

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Printf("Error getting database driver: %v", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	if err != nil {
		fmt.Printf("NewWithDatabaseInstance: %v\n", err)
		return err
	}

	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	return nil
}
