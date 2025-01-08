package server

import (
	"net/http"
	"os"
	"shinko/internal/handlers"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func StartApp(address string) {

	godotenv.Load()

	serveMux := http.NewServeMux()

	apiConfig := &handlers.ApiConfig{
		DbQueries: setupDatabase(),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	RegisterHandlers(serveMux, apiConfig)

	server := http.Server{
		Addr:    address,
		Handler: serveMux,
	}

	server.ListenAndServe()
}
