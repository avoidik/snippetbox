package main

import "snippetbox.org/pkg/models"

type App struct {
	htmlDir   string
	addr      string
	staticDir string
	database  *models.Database
}
