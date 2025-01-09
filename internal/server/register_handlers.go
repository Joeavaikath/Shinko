package server

import (
	"net/http"
	"shinko/internal/handlers"
)

func RegisterHandlers(s *http.ServeMux, apiConfig *handlers.ApiConfig) {

	handlers := []func(*http.ServeMux, *handlers.ApiConfig){
		handlers.UserRoutes,
		handlers.AdminRoutes,
		handlers.MetricsRoutes,
		handlers.TokenRoutes,
		handlers.ActionRoutes,
	}

	for _, handler := range handlers {
		handler(s, apiConfig)
	}

}
