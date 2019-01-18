package main

import (
	"log"
	"net/http"
	"os"

	"github.com/leighmcculloch/looks.wtf/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(service.Handler)))
}
