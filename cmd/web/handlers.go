package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.RenderHtml(w, "home.page.html", nil)
}

func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")
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

	app.RenderHtml(w, "show.page.html", &HtmlData{Snippet: snippet})
}

func (app *App) LatestSnippets(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.database.LatestSnippets()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	if snippets == nil {
		app.NotFound(w)
		return
	}

	for _, s := range snippets {
		app.RenderHtml(w, "show.page.html", &HtmlData{Snippet: s})
	}
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("NewSnippet"))
	case http.MethodPost:
		id, err := app.database.InsertSnippet("test", "test", "+2 days")
		if err != nil {
			app.ServerError(w, err)
			return
		}
		fmt.Fprintf(w, "New item id = %d\n", id)
	default:
		app.ClientError(w, http.StatusNotImplemented)
	}
}

func (app *App) VersionInfo(w http.ResponseWriter, r *http.Request) {
	verFile := filepath.Join(app.staticDir, "VERSION")
	if _, err := os.Stat(verFile); err != nil {
		http.Error(w, "Version was not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, verFile)
}
