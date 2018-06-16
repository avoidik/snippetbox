package main

import (
	"flag"
	"log"
	"net/http"
)

type Config struct {
	addr      string
	staticDir string
}

func main() {
	log.Println("Starting the server...")

	cfg := new(Config)
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address to listen")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/version", VersionInfo)

	log.Printf("Listening on %s\n", cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
