package server

import (
	"log"
	"net/http"
	"os"
	"shinko/internal/handlers"
	"time"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func StartApp(address string) {

	// Only used in local, k8s will load into OS
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file, are you on non-local?: %v", err)
	}

	serveMux := http.NewServeMux()

	apiConfig := &handlers.ApiConfig{
		DbQueries: setupDatabase(),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	RegisterHandlers(serveMux, apiConfig)

	server := http.Server{
		Addr:              address,
		Handler:           serveMux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not listen on %s: %v", address, err)
	}
}
