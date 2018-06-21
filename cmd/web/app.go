package main

import (
	"net/http"

	"github.com/alexedwards/scs"
	"snippetbox.org/pkg/models"
)

type App struct {
	htmlDir      string
	addr         string
	staticDir    string
	databaseFile string
	secret       string
	tlsCert      string
	tlsKey       string
	server       *http.Server
	sessions     *scs.Manager
	database     *models.Database
}
