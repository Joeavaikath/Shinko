package handlers

import (
	"errors"
	"net/http"
	"shinko/internal/auth"
	"shinko/util"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) GetBearerAndValidate(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {

	accessToken, err := auth.GetBearerToken(r.Header)

	if accessToken == "" {
		util.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return uuid.Nil, errors.New("empty auth token")
	}

	if util.ErrorNotNil(err, w) {
		return uuid.Nil, errors.New("error parsing auth token")
	}

	userID, err := auth.ValidateJWT(accessToken, cfg.JwtSecret)
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return uuid.Nil, errors.New("jwt validation failed")
	}

	return userID, nil
}
