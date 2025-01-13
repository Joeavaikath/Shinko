package handlers

import (
	"database/sql"
	"net/http"
	"shinko/internal/database"
	"shinko/util"

	"github.com/google/uuid"
)

func EventRoutes(s *http.ServeMux, apiConfig *ApiConfig) {
	s.Handle("POST /api/events", http.HandlerFunc(apiConfig.addEvent))
}

func (cfg *ApiConfig) addEvent(w http.ResponseWriter, r *http.Request) {

	type addEventParams struct {
		Action_id string `json:"action_id"`
		Comment   string `json:"comment"`
	}

	user_uuid, err := cfg.GetBearerAndValidate(w, r)
	if err != nil {
		return
	}

	// Decode params from incoming request
	params, err := util.DecodeJSON[addEventParams](r)
	if err != nil {
		return
	}

	action_uuid, err := uuid.Parse(params.Action_id)
	if err != nil {
		return
	}
	createEventDbParams := database.CreateEventParams{
		ActionID: action_uuid,
		UserID:   user_uuid,
		Comment:  sql.NullString{String: params.Comment, Valid: params.Comment != ""},
	}

	event, err := cfg.DbQueries.CreateEvent(r.Context(), createEventDbParams)
	if err != nil {
		return
	}

	type eventCreatedResponse struct {
		Id      uuid.UUID `json:"event_id"`
		Comment string    `json:"comment"`
	}

	eventResponseParams := eventCreatedResponse{
		event.ID,
		event.Comment.String,
	}

	// If success. err is nil so defer does nothing.
	util.RespondWithJSON(w, http.StatusCreated, eventResponseParams)

}
