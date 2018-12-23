package main

import (
	"log"
	"net/http"
	"os"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (a appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := a(w, r); err != nil {
		log.Printf("Error: %#v", err)
		http.Error(w, "There was an error. Please try again.", 500)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	http.Handle("/oauth", appHandler(oauthHandler))
	http.Handle("/command/look", appHandler(commandLookHandler))
	http.Handle("/command/looks", appHandler(commandLooksHandler))
	http.Handle("/action", appHandler(actionHandler))

	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
