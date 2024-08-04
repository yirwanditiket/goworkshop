package main

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := doThis()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("handling %q: %v", r.RequestURI, err)
		return
	}

	err = doThat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("handling %q: %v", r.RequestURI, err)
		return
	}
}

func doThis() error { return nil }

func doThat() error { return nil }
