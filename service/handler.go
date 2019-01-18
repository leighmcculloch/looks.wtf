package service

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

var mux = func() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/oauth", appHandler(oauthHandler))
	mux.Handle("/command/look", appHandler(commandLookHandler(dataLooks, dataTags)))
	mux.Handle("/command/looks", appHandler(commandLooksHandler(dataLooks, dataTags)))
	mux.Handle("/action", appHandler(actionHandler))
	return mux
}()

func Handler(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}
