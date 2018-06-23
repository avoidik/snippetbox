package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/justinas/nosurf"
	"snippetbox.org/pkg/models"
)

type HtmlData struct {
	CSRFtoken string
	Form      interface{}
	Path      string
	Flash     string
	LoggedIn  bool
	Snippet   *models.Snippet
	Snippets  []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format(time.Stamp)
}

func (app *App) RenderHtml(w http.ResponseWriter, r *http.Request, page string, data *HtmlData) {
	if data == nil {
		data = &HtmlData{}
	}

	data.Path = r.URL.Path

	data.CSRFtoken = nosurf.Token(r)

	var err error
	data.LoggedIn, err = app.LoggedIn(r)
	if err != nil {
		app.ServerError(w, err)
		return
	}

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
