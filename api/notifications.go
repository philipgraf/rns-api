package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/philipgraf/rns-api/database"
	"github.com/philipgraf/rns-api/models"
)

func indexNotiHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetFromContext(r)

	var result = []models.Noti{}
	offset := 0

	const PAGE_LIMIT int = 50

	queryParams := r.URL.Query()

	page, err := strconv.Atoi(queryParams.Get("page"))

	if err != nil {
		page = 0
	}

	if page > 0 {
		offset = (page - 1) * PAGE_LIMIT
	}

	if err := db.C("notifications").Find(nil).Skip(offset).Limit(PAGE_LIMIT).All(&result); err != nil {
		log.Fatalf("unable to get jan: %v", err)
	}

	if len(result) == PAGE_LIMIT {
		nextPage := strconv.Itoa(page + 1)
		// TODO get whole url
		w.Header().Add("Link", "</?page="+nextPage+"; rel=\"next\">")
	}

	if page > 0 {
		prevPage := strconv.Itoa(page - 1)
		// TODO get whole url
		w.Header().Add("Link", "</?page="+prevPage+"; rel=\"prev\">")
	}

	writeJson(w, result)
}

func createNotiHandler(w http.ResponseWriter, r *http.Request) {

	noti := &models.Noti{}
	db := database.GetFromContext(r)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(noti)
	if err != nil {
		Error(w, err, http.StatusBadRequest)
	}

	if err := db.C("notifications").Insert(noti); err != nil {
		log.Fatalf("unable to insert into user: %v", err)
		Error(w, err, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
