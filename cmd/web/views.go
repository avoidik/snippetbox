package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"snippetbox.org/pkg/models"
)

type HtmlData struct {
	Snippet *models.Snippet
}

func (app *App) RenderHtml(w http.ResponseWriter, page string, data *HtmlData) {
	files := []string{
		filepath.Join(app.htmlDir, "base.html"),
		filepath.Join(app.htmlDir, page),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}
}
