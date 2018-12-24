package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/leighmcculloch/looks.wtf/service/shared/looks"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (a appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := a(w, r); err != nil {
		log.Printf("Error: %#v", err)
		http.Error(w, "There was an error. Please try again.", 500)
	}
}

var dataLooks = looks.ParseLooks(func() io.Reader {
	f, err := os.Open("data/looks.yml")
	if err != nil {
		panic(err)
	}
	return f
}())

var dataTags = looks.ParseTags(func() io.Reader {
	f, err := os.Open("data/tags.yml")
	if err != nil {
		panic(err)
	}
	return f
}())

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	http.Handle("/oauth", appHandler(oauthHandler))
	http.Handle("/command/look", appHandler(commandLookHandler(dataLooks, dataTags)))
	http.Handle("/command/looks", appHandler(commandLooksHandler(dataLooks, dataTags)))
	http.Handle("/action", appHandler(actionHandler))

	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
