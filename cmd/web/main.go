package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func existDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	log.Println("Starting the server...")

	app := new(App)
	flag.StringVar(&app.addr, "addr", ":4000", "HTTP network address to listen")
	flag.StringVar(&app.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&app.htmlDir, "html-dir", "./ui/html", "Path to html templates")
	flag.Parse()

	if !existDir(app.staticDir) {
		log.Fatal("Folder for static-dir was not found")
	}

	if !existDir(app.htmlDir) {
		log.Fatal("Folder for html-dir was not found")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	fileServer := http.FileServer(http.Dir(app.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/version", app.VersionInfo)

	log.Printf("Listening on %s\n", app.addr)
	err := http.ListenAndServe(app.addr, mux)
	log.Fatal(err)
}
