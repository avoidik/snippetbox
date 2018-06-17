package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

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
	flag.Parse()

	if !existDir(app.staticDir) {
		log.Fatal("Folder for static-dir was not found")
	}

	if !existDir(app.htmlDir) {
		log.Fatal("Folder for html-dir was not found")
	}

	server := &http.Server{
		Addr:    app.addr,
		Handler: app.Routes(),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Terminating (signal caught - %s)\n", sig)
			server.Shutdown(context.Background())
		}
	}()

	db := connect()
	defer func() {
		log.Println("Closing database connection")
		db.Close()
	}()

	app.database = &models.Database{DB: db}

	log.Printf("Listening on %s\n", app.addr)

	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Println("The error below raised after shutdown:")
			log.Println(err)
		}
		return
	}
}

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "file:info.db?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
