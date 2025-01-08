package server

import (
	"database/sql"
	"os"
	"shinko/internal/database"
)

func setupDatabase() *database.Queries {
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		os.Exit(1)
	}

	return database.New(db)
}
