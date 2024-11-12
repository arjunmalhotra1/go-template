package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "host=postgres port=5432 user=user password=password dbname=db sslmode=disable"
	// Note the host should be the container image name used in the docker compose.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "error in sql.Open")
	}

	//Run migrations
	if err := RunMigrations(db); err != nil {
		return nil, errors.Wrap(err, "failed to run migrations")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "error in database ping")
	}
	fmt.Println("Database connected successfully!")
	return db, nil
}
