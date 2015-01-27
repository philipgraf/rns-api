package api

import (
	"encoding/json"
	"net/http"

	"github.com/philipgraf/rns-api/database"
)

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func connectDB(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := database.Connect()
		if err != nil {
			Error(w, err, http.StatusInternalServerError)
			return
		}
		database.SetToContext(r, db)
		h.ServeHTTP(w, r)
	}
}

func writeJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}
