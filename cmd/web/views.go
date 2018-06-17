package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (app *App) RenderHtml(w http.ResponseWriter, page string) {
	files := []string{
		filepath.Join(app.htmlDir, "base.html"),
		filepath.Join(app.htmlDir, page),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.ServerError(w, err)
		return
	}
}
