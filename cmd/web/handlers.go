package main

import (
	"fmt"
	"html/template"
	"log"
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

	files := []string{
		filepath.Join(app.htmlDir, "base.html"),
		filepath.Join(app.htmlDir, "home.page.html"),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryID)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "ShowSnippet with %d id\n", id)
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NewSnippet"))
}

func (app *App) VersionInfo(w http.ResponseWriter, r *http.Request) {
	verFile := filepath.Join(app.staticDir, "VERSION")
	if _, err := os.Stat(verFile); err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Version was not found"))
		return
	}
	http.ServeFile(w, r, verFile)
}
