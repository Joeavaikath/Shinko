package config

import (
	"shinko/internal/database"
	"sync/atomic"
)

type ApiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
	jwtSecret      string
	polkaKey       string
}

func NewApiConfig(dbQueries *database.Queries) *ApiConfig {
	return &ApiConfig{
		dbQueries: dbQueries,
	}
}
