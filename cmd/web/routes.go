package main

import (
	"net/http"
)

func (app *App) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)
	mux.HandleFunc("/latest", app.LatestSnippets)

	fileServer := http.FileServer(http.Dir(app.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/version", app.VersionInfo)
	return mux
}
