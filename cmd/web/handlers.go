package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Write([]byte("Hello From Snippetbox"))
}

// snippetView is an HTTP handler function that displays a specific snippet
// based on the "id" query parameter in the request URL. If the "id" parameter
// is invalid or less than 1, the function will return a 404 Not Found response.
func snippetView(w http.ResponseWriter, req *http.Request) {
	id, error := strconv.Atoi(req.URL.Query().Get("id"))
	if error != nil || id < 1 {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("In Snippetbox create"))
}
