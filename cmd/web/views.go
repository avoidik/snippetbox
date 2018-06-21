package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"snippetbox.org/pkg/models"
)

type HtmlData struct {
	Form     interface{}
	Path     string
	Flash    string
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format(time.Stamp)
}

func (app *App) RenderHtml(w http.ResponseWriter, r *http.Request, page string, data *HtmlData) {
	if data == nil {
		data = &HtmlData{}
	}

	data.Path = r.URL.Path

	files := []string{
		filepath.Join(app.htmlDir, "base.html"),
		filepath.Join(app.htmlDir, page),
	}

	fn := template.FuncMap{
		"humanDate": humanDate,
	}

	ts, err := template.New("").Funcs(fn).ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)
}
