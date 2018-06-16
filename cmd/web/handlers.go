package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ShowSnippet"))
}

func NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NewSnippet"))
}
