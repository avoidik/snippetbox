package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.Home))
	mux.Get("/snippet/new", http.HandlerFunc(app.NewSnippet))
	mux.Post("/snippet/new", http.HandlerFunc(app.CreateSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.ShowSnippet))

	fileServer := http.FileServer(http.Dir(app.staticDir))
	mux.Get("/static/", http.StripPrefix("/static", DisableIndex(fileServer)))

	mux.Get("/version", http.HandlerFunc(app.VersionInfo))
	return mux
}
