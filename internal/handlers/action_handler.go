package handlers

import (
	"database/sql"
	"net/http"
	"shinko/internal/database"
	"shinko/util"

	"github.com/google/uuid"
)

func ActionRoutes(s *http.ServeMux, apiConfig *ApiConfig) {
	s.Handle("POST /api/actions", http.HandlerFunc(apiConfig.addAction))
}

func (cfg *ApiConfig) addAction(w http.ResponseWriter, r *http.Request) {

	// Define request param structure
	type addActionParams struct {
		UserId      string `json:"user_id"`
		Action_name string `json:"action_name"`
		Description string `json:"description"`
	}

	userID, err := cfg.GetBearerAndValidate(w, r)
	if err != nil {
		return
	}

	// Decode params from incoming request
	params, err := util.DecodeJSON[addActionParams](r)
	if err != nil {
		return
	}

	createActionDbParams := database.CreateActionParams{
		UserID:      userID,
		Name:        params.Action_name,
		Description: sql.NullString{String: params.Description, Valid: params.Description != ""},
	}

	action, err := cfg.DbQueries.CreateAction(r.Context(), createActionDbParams)

	type actionCreatedResponse struct {
		Name        string    `json:"name"`
		Id          uuid.UUID `json:"action_id"`
		UserID      uuid.UUID `json:"user_id"`
		Description string    `json:"description"`
	}

	actionResponseParams := actionCreatedResponse{
		Name:        action.Name,
		Description: action.Description.String,
		Id:          action.ID,
		UserID:      action.UserID,
	}

	// If success. err is nil so defer does nothing.
	util.RespondWithJSON(w, http.StatusCreated, actionResponseParams)

}
