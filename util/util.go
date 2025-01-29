package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func DecodeJSON[T any](r *http.Request) (T, error) {
	var result T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)
	return result, err
}

func ErrorNotNil(err error, w http.ResponseWriter) bool {
	if err != nil {
		RespondWithError(w, 500, error.Error(err))
		return true
	}
	return false
}

func RespondWithError(w http.ResponseWriter, code int, errorPayload interface{}) {
	w.WriteHeader(code)
	dat, _ := json.Marshal(errorPayload)
	_, err := w.Write(dat)
	if err != nil {
		log.Printf("Failed to write error response: %v", err)
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	dat, _ := json.Marshal(payload)
	_, err := w.Write(dat)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func SliceContains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
