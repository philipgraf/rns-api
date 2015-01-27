package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philipgraf/rns-api/config"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	notiRouter := r.PathPrefix("/notis").Subrouter()
	notiRouter.HandleFunc("/", use(indexNotiHandler, connectDB))
	return r
}

func Start() error {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("Unable to read config: %v", err)
		return err
	}
	fmt.Printf("listen on %v", conf.Addr)
	return http.ListenAndServe(conf.Addr, Router())
}
