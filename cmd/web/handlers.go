package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"snippetbox.org/pkg/forms"
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

	session := app.sessions.Load(r)
	flash, err := session.PopString(w, "flash")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHtml(w, r, "show.page.html", &HtmlData{
		Snippet: snippet,
		Flash:   flash,
	})
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	app.RenderHtml(w, r, "new.page.html", &HtmlData{
		Form: &forms.NewSnippet{},
	})
}

func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.NewSnippet{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: r.PostForm.Get("expires"),
	}

	if !form.Valid() {
		app.RenderHtml(w, r, "new.page.html", &HtmlData{Form: form})
		return
	}

	formatDate := fmt.Sprintf("+%s.0 seconds", form.Expires)

	id, err := app.database.InsertSnippet(form.Title, form.Content, formatDate)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	session := app.sessions.Load(r)
	err = session.PutString(w, "flash", "Snipped saved succesfully")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *App) VersionInfo(w http.ResponseWriter, r *http.Request) {
	verFile := filepath.Join(app.staticDir, "VERSION")
	if _, err := os.Stat(verFile); err != nil {
		http.Error(w, "Version was not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, verFile)
}

func DisableIndex(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.Error(w, "Nothing to see here", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *App) SignupUser(w http.ResponseWriter, r *http.Request) {
	app.RenderHtml(w, r, "signup.page.html", &HtmlData{
		Form: &forms.SignupUser{},
	})
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.SignupUser{
		Name:     r.PostForm.Get("name"),
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	if !form.Valid() {
		app.RenderHtml(w, r, "signup.page.html", &HtmlData{Form: form})
		return
	}

	fmt.Fprintf(w, "CreateUser")
}

func (app *App) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "LoginUser")
}

func (app *App) VerifyUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "VerifyUser")
}

func (app *App) LogoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "LogoutUser")
}
