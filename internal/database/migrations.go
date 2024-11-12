package database

import (
	"database/sql"
	"fmt"
	"log"
)

func RunMigrations(db *sql.DB) error {
	migrations := []string{
		// Create the tasks table
		`
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			completed BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);
		`,
		// Add other migrations here if needed
	}

	for _, query := range migrations {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error running migration: %w", err)
		}
		log.Println("Migration executed successfully")
	}
	return nil
}
