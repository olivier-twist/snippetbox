package main

import (
	"log"
	"net/http"
)

func main() {
	mutex := http.NewServeMux()
	mutex.HandleFunc("/", home)
	mutex.HandleFunc("/snippet/view", snippetView)
	mutex.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting Server")
	error := http.ListenAndServe(":4000", mutex)
	log.Fatal(error)
}
