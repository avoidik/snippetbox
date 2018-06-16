package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
