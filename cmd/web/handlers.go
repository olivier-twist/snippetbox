package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Home takes the user to the home page
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Invalid Path\n", http.StatusNotFound)
		return
	}

	ts, err := template.ParseFiles("./ui/html/base.tmpl", "./ui/html/pages/home.tmpl")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// CreateSnippet creates a snippet.
func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Add("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create Snippet\n"))
}

// View Snippet views a snippet.
func ViewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Snippet Create\n"))
}
