package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.database.LatestSnippets()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHtml(w, r, "home.page.html", &HtmlData{Snippets: snippets})
}

func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get(":id")
	id, err := strconv.Atoi(queryID)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.database.GetSnippet(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	if snippet == nil {
		app.NotFound(w)
		return
	}

	app.RenderHtml(w, r, "show.page.html", &HtmlData{Snippet: snippet})
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NewSnippet"))
}

func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateSnippet"))
}

func (app *App) VersionInfo(w http.ResponseWriter, r *http.Request) {
	verFile := filepath.Join(app.staticDir, "VERSION")
	if _, err := os.Stat(verFile); err != nil {
		http.Error(w, "Version was not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, verFile)
}
