package handlers

import (
	"net/http"
	"shinko/internal/auth"
	"shinko/internal/database"
	"shinko/internal/models"
	"shinko/util"

	"github.com/google/uuid"
)

func UserRoutes(s *http.ServeMux, apiConfig *ApiConfig) {
	s.Handle("POST /api/users", http.HandlerFunc(apiConfig.addUser))
	s.Handle("PUT /api/users", http.HandlerFunc(apiConfig.updateUser))

	s.Handle("GET /api/users/{user_id}/actions", http.HandlerFunc(apiConfig.getUserActions))
}

func (cfg *ApiConfig) addUser(w http.ResponseWriter, r *http.Request) {

	type addUserRequestParams struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	params, err := util.DecodeJSON[addUserRequestParams](r)
	if util.ErrorNotNil(err, w) {
		return
	}

	hashed_password, err := auth.HashPassword(params.Password)
	if util.ErrorNotNil(err, w) {
		return
	}

	createUserParams := database.CreateUserParams{
		Email:        params.Email,
		Username:     params.Username,
		PasswordHash: hashed_password,
	}

	user, err := cfg.DbQueries.CreateUser(r.Context(), createUserParams)
	if util.ErrorNotNil(err, w) {
		return
	}

	userCreatedResponse := models.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     params.Email,
		Username:  params.Username,
	}

	util.RespondWithJSON(w, 201, userCreatedResponse)
}

func (cfg *ApiConfig) updateUser(w http.ResponseWriter, r *http.Request) {

	type updateUserRequestParams struct {
		Password string `json:"password"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	userID, err := cfg.GetBearerAndValidate(w, r)

	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	params, err := util.DecodeJSON[updateUserRequestParams](r)
	if util.ErrorNotNil(err, w) {
		return
	}

	hashedPass, err := auth.HashPassword(params.Password)
	if util.ErrorNotNil(err, w) {
		return
	}

	dbParams := database.UpdateUserParams{
		PasswordHash: hashedPass,
		Email:        params.Email,
		Username:     params.Username,
		ID:           userID,
	}

	err = cfg.DbQueries.UpdateUser(r.Context(), dbParams)
	if util.ErrorNotNil(err, w) {
		return
	}

	user, err := cfg.DbQueries.GetUserByEmail(r.Context(), dbParams.Email)
	if util.ErrorNotNil(err, w) {
		return
	}

	userResponse := models.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Username:  user.Username,
	}

	util.RespondWithJSON(w, 200, userResponse)

}

func (cfg *ApiConfig) getUserActions(w http.ResponseWriter, r *http.Request) {

	user_uuid, err := cfg.GetBearerAndValidate(w, r)
	if err != nil {
		return
	}

	user_id := r.PathValue("user_id")
	path_user_uuid, err := uuid.Parse(user_id)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	if path_user_uuid != user_uuid {
		util.RespondWithError(w, http.StatusUnauthorized, struct {
			Error string `json:"error"`
		}{Error: "unauthorized user operation"})
		return
	}

	actions, err := cfg.DbQueries.GetUserActions(r.Context(), user_uuid)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, actions)
}
