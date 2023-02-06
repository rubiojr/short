// Package handlers provides HTTP request handlers.
package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rubiojr/short/storages"
)

func EncodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if url := r.PostFormValue("url"); url != "" {
			code, err := storage.Save(url)
			if err != nil {
				return
			}
			rurl := fmt.Sprintf("%s/red/%s", os.Getenv("SHORT_BASE_URL"), code)
			w.Write([]byte(rurl))
		}
	}

	return http.HandlerFunc(handleFunc)
}

func DecodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/dec/"):]

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
			return
		}

		w.Write([]byte(url))
	}

	return http.HandlerFunc(handleFunc)
}

func RedirectHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/red/"):]

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
			return
		}

		http.Redirect(w, r, string(url), 301)
	}

	return http.HandlerFunc(handleFunc)
}
