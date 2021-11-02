package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/leighmcculloch/looks.wtf/service/data"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/oauth", appHandler(oauthHandler))
	mux.Handle("/command/look", appHandler(commandLookHandler(data.Looks, data.Tags)))
	mux.Handle("/command/looks", appHandler(commandLooksHandler(data.Looks, data.Tags)))
	mux.Handle("/action", appHandler(actionHandler))
	mux.Handle("/", http.FileServer(http.FS(staticSub)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

//go:embed static
var static embed.FS

var staticSub = func() fs.FS {
	fs, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	return fs
}()

type appHandler func(http.ResponseWriter, *http.Request) error

func (a appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := a(w, r); err != nil {
		log.Printf("Error: %#v", err)
		http.Error(w, "There was an error. Please try again.", 500)
	}
}
