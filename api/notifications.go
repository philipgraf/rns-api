package api

import (
	"fmt"
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

	page, err := strconv.Atoi(queryParams["page"][0])

	if err != nil {
		page = 0
	}

	if page > 0 {
		offset = (page - 1) * PAGE_LIMIT
	}

	if err := db.C("notifications").Find(nil).Skip(offset).Limit(PAGE_LIMIT).All(&result); err != nil {
		log.Fatalf("unable to get jan: %v", err)
	}

	fmt.Println(r.URL.Scheme)

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
