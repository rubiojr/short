package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rubiojr/short/handlers"
	"github.com/rubiojr/short/storages"
)

func main() {
	storage := &storages.Sqlite{}
	err := storage.Init(".")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handlers.EncodeHandler(storage))
	http.Handle("/dec/", handlers.DecodeHandler(storage))
	http.Handle("/red/", handlers.RedirectHandler(storage))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
