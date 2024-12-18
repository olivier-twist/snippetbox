package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// home is an HTTP handler function that displays the home page. If the request
// URL path is not "/", the function will return a 404 Not Found response.
func home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

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
	fmt.Fprintf(w, "Display a specific snippet with ID %d\n", id)
}

// snippetCreate is an HTTP handler function that creates a new snippet. It
// expects a POST request with a form-encoded body containing the title and
// content of the new snippet. If the request method is not POST, the function
// will return a 405 Method Not Allowed response. If the request is successful,
// the function will write a simple confirmation message to the response.
func snippetCreate(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("In Snippetbox create\n"))
}
