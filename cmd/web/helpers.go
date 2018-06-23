package main

import (
	"net/http"
	"os"
)

func (app *App) LoggedIn(r *http.Request) (bool, error) {
	session := app.sessions.Load(r)
	loggedIn, err := session.Exists("currentUserId")
	if err != nil {
		return false, err
	}
	return loggedIn, nil
}

func existDir(path *string) bool {
	if _, err := os.Stat(*path); os.IsNotExist(err) {
		return false
	}
	return true
}
