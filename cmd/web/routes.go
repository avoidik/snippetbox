package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", alice.New(NoSurf).Then(http.HandlerFunc(app.Home)))

	mux.Get("/snippet/new", alice.New(app.RequireLogin, NoSurf).Then(http.HandlerFunc(app.NewSnippet)))
	mux.Post("/snippet/new", alice.New(app.RequireLogin, NoSurf).Then(http.HandlerFunc(app.CreateSnippet)))
	mux.Get("/snippet/:id", alice.New(NoSurf).Then(http.HandlerFunc(app.ShowSnippet)))

	mux.Get("/user/signup", alice.New(NoSurf).Then(http.HandlerFunc(app.SignupUser)))
	mux.Post("/user/signup", alice.New(NoSurf).Then(http.HandlerFunc(app.CreateUser)))
	mux.Get("/user/login", alice.New(NoSurf).Then(http.HandlerFunc(app.LoginUser)))
	mux.Post("/user/login", alice.New(NoSurf).Then(http.HandlerFunc(app.VerifyUser)))
	mux.Post("/user/logout", alice.New(app.RequireLogin, NoSurf).Then(http.HandlerFunc(app.LogoutUser)))

	fileServer := http.FileServer(http.Dir(app.staticDir))
	mux.Get("/static/", alice.New(StripStatic, DisableIndex).Then(fileServer))

	mux.Get("/version", http.HandlerFunc(app.VersionInfo))

	return alice.New(LogRequest, SecureHeaders).Then(mux)
}
