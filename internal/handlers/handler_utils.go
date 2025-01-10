package handlers

import (
	"net/http"
	"shinko/internal/auth"
	"shinko/util"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) GetBearerAndValidate(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {

	accessToken, err := auth.GetBearerToken(r.Header)

	if accessToken == "" {
		util.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return uuid.Nil, err
	}

	if util.ErrorNotNil(err, w) {
		return uuid.Nil, err
	}

	userID, err := auth.ValidateJWT(accessToken, cfg.JwtSecret)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return uuid.Nil, err
	}

	return userID, nil
}
