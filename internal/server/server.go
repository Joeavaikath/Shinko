package server

import (
	"database/sql"
	"net/http"
	"os"
	"shinko/internal/database"
	"shinko/internal/handlers"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func StartApp(address string) {

	godotenv.Load()

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		os.Exit(1)
	}
	dbQueries := database.New(db)

	serveMux := http.NewServeMux()

	apiConfig := &handlers.ApiConfig{
		DbQueries: dbQueries,
		JwtSecret: os.Getenv("SECRET"),
	}

	RegisterHandlers(serveMux, apiConfig)

	server := http.Server{
		Addr:    address,
		Handler: serveMux,
	}

	server.ListenAndServe()
}
