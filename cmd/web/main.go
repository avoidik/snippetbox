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

	log.Printf("Listening on %s\n", app.addr)
	err := http.ListenAndServe(app.addr, app.Routes())
	log.Fatal(err)
}
