package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
