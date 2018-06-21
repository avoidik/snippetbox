package main

import "net/http"

func (app *App) LoggedIn(r *http.Request) (bool, error) {
	session := app.sessions.Load(r)
	loggedIn, err := session.Exists("currentUserId")
	if err != nil {
		return false, err
	}
	return loggedIn, nil
}
