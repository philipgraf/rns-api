package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/philipgraf/rns-api/api"
)

func main() {

	go startAPI()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	<-sigc
}

func startAPI() {
	if err := api.Start(); err != nil {
		log.Fatalf("API stopped: %v", err)
	}
}
