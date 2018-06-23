package main

import (
	"flag"
	"log"
	"time"

	"github.com/alexedwards/scs"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Starting the server...")

	app := new(App)
	flag.StringVar(&app.addr, "addr", ":4000", "HTTP network address to listen")
	flag.StringVar(&app.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&app.htmlDir, "html-dir", "./ui/html", "Path to html templates")
	flag.StringVar(&app.databaseFile, "db-file", "./info.db", "Path to database file")
	flag.StringVar(&app.secret, "secret", "8sB9ozuKkqWtN3b6lEiInd1dSISxPWogpaGV5HG4wKs=", "Secret key for cookies encryption")
	flag.StringVar(&app.tlsCert, "tls-cert", "./tls/cert.pem", "TLS certificate")
	flag.StringVar(&app.tlsKey, "tls-key", "./tls/key.pem", "TLS private-key")
	flag.Parse()

	if !existDir(&app.staticDir) {
		log.Fatal("Folder for static-dir was not found")
	}

	if !existDir(&app.htmlDir) {
		log.Fatal("Folder for html-dir was not found")
	}

	if !existDir(&app.tlsCert) {
		log.Fatal("TLS certificate was not found")
	}

	if !existDir(&app.tlsKey) {
		log.Fatal("TLS key was not found")
	}

	if err := app.ConnectDb(); err != nil {
		log.Fatal("Failed to establish database connection")
	}
	defer app.CloseDB()

	sessionManager := scs.NewCookieManager(app.secret)
	sessionManager.Lifetime(12 * time.Hour)
	sessionManager.Persist(true)
	sessionManager.Secure(true)
	app.sessions = sessionManager

	app.InitServer()

	app.MonitorInterrupts()

	app.RunServer()
}
