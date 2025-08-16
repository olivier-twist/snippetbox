package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippetbox/create", CreateSnippet)
	mux.HandleFunc("/snippetbox/view", ViewSnippet)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
