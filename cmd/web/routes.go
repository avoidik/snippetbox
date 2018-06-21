package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.Home))

	mux.Get("/snippet/new", http.HandlerFunc(app.NewSnippet))
	mux.Post("/snippet/new", http.HandlerFunc(app.CreateSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.ShowSnippet))

	mux.Get("/user/signup", http.HandlerFunc(app.SignupUser))
	mux.Post("/user/signup", http.HandlerFunc(app.CreateUser))
	mux.Get("/user/login", http.HandlerFunc(app.LoginUser))
	mux.Post("/user/login", http.HandlerFunc(app.VerifyUser))
	mux.Post("/user/logout", http.HandlerFunc(app.LogoutUser))

	fileServer := http.FileServer(http.Dir(app.staticDir))
	mux.Get("/static/", http.StripPrefix("/static", DisableIndex(fileServer)))

	mux.Get("/version", http.HandlerFunc(app.VersionInfo))

	chain := alice.New(LogRequest, SecureHeaders).Then(mux)

	return chain
}
