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

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/version", VersionInfo)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
