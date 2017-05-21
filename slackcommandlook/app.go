package slackcommandlook

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type appHandler func(context.Context, http.ResponseWriter, *http.Request) error

func (a appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := a(c, w, r); err != nil {
		log.Errorf(c, "Error: %#v", err)
		http.Error(w, "There was an error. Please try again.", 500)
	}
}

func init() {
	http.Handle("/command/look", appHandler(commandLookHandler))
}
