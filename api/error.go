package api

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func Error(w http.ResponseWriter, err error, status int) {
	response := errorResponse{err.Error()}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	if err = enc.Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
