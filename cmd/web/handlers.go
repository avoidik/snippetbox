package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryID)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "ShowSnippet with %d id\n", id)
}

func NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NewSnippet"))
}
