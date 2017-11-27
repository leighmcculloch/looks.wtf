package slackcommandlook

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (a appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	r = r.WithContext(c)
	if err := a(w, r); err != nil {
		log.Errorf(c, "Error: %#v", err)
		http.Error(w, "There was an error. Please try again.", 500)
	}
}

func init() {
	http.Handle("/command/look", appHandler(commandLookHandler))
}
