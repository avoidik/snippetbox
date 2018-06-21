package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexedwards/scs"

	_ "github.com/mattn/go-sqlite3"
	"snippetbox.org/pkg/models"
)

func existDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	log.Println("Starting the server...")

	app := new(App)
	flag.StringVar(&app.addr, "addr", ":4000", "HTTP network address to listen")
	flag.StringVar(&app.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&app.htmlDir, "html-dir", "./ui/html", "Path to html templates")
	flag.StringVar(&app.databaseFile, "db-file", "./info.db", "Path to database file")
	flag.StringVar(&app.secret, "secret", "8sB9ozuKkqWtN3b6lEiInd1dSISxPWogpaGV5HG4wKs=", "Secret key for cookies encryption")
	tlsCert := flag.String("tls-cert", "./tls/cert.pem", "TLS certificate")
	tlsKey := flag.String("tls-key", "./tls/key.pem", "TLS private-key")
	flag.Parse()

	if !existDir(app.staticDir) {
		log.Fatal("Folder for static-dir was not found")
	}

	if !existDir(app.htmlDir) {
		log.Fatal("Folder for html-dir was not found")
	}

	if !existDir(*tlsCert) {
		log.Fatal("TLS certificate was not found")
	}

	if !existDir(*tlsKey) {
		log.Fatal("TLS key was not found")
	}

	if err := app.connectDb(); err != nil {
		log.Fatal("Failed to establish database connection")
	}
	defer func() {
		log.Println("Closing database connection")
		app.closeDB()
	}()

	sessionManager := scs.NewCookieManager(app.secret)
	sessionManager.Lifetime(12 * time.Hour)
	sessionManager.Persist(true)
	sessionManager.Secure(true)
	app.sessions = sessionManager

	server := &http.Server{
		Addr:         app.addr,
		Handler:      app.Routes(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Terminating (signal caught - %s)\n", sig)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer func() {
				log.Println("Closing context")
				cancel()
			}()
			server.Shutdown(ctx)
		}
	}()

	log.Printf("Listening on %s\n", app.addr)

	if err := server.ListenAndServeTLS(*tlsCert, *tlsKey); err != nil {
		if err != http.ErrServerClosed {
			log.Println("The error below raised after shutdown:")
			log.Println(err)
		}
	}
}

func (app *App) connectDb() error {
	initDb := !existDir(app.databaseFile)

	dsn := fmt.Sprintf("file:%s?cache=shared&_loc=auto", app.databaseFile)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			db.Close()
		}
	}()

	if err = db.Ping(); err != nil {
		return err
	}

	app.database = &models.Database{DB: db}

	if initDb {
		log.Println("Initializing database...")
		if err := app.database.InitializeDb(); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) closeDB() {
	if app.database != nil {
		app.database.Close()
	}
}
