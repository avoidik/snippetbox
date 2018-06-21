package main

import (
	"github.com/alexedwards/scs"
	"snippetbox.org/pkg/models"
)

type App struct {
	htmlDir      string
	addr         string
	staticDir    string
	databaseFile string
	secret       string
	sessions     *scs.Manager
	database     *models.Database
}
