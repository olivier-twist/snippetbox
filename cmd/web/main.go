package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippetbox/create", CreateSnippet)
	mux.HandleFunc("/snippetbox/view", ViewSnippet)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
