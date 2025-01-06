package main

import (
	"database/sql"
	"net/http"
	"os"
	"shinko/internal/config"
	"shinko/internal/database"
	"shinko/internal/logger"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		os.Exit(1)
	}

	dbQueries := database.New(db)

	apiConfig := config.NewApiConfig(dbQueries)

	serveMux := http.NewServeMux()

	configureRoutes(serveMux, apiConfig)

	server := http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	server.ListenAndServe()

}

func configureRoutes(mainMux *http.ServeMux, apiConfig *config.ApiConfig) {

	logger.Info(apiConfig)
	mainMux.Handle("GET /users", http.RedirectHandler("/userz", http.StatusFound))

}
