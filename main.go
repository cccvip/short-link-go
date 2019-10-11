package main

import (
	"github.com/carl-xiao/short-link-go/api"
	"log"
	"net/http"
	"time"
)

func main() {
	app := api.App{}
	app.Initliaze()
	srv := &http.Server{
		Handler:      app.Router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("server start port %s", "8000")
	log.Fatal(srv.ListenAndServe())
}
