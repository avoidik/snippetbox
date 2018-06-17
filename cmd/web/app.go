package main

import "snippetbox.org/pkg/models"

type App struct {
	htmlDir      string
	addr         string
	staticDir    string
	databaseFile string
	database     *models.Database
}
