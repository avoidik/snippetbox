package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"snippetbox.org/pkg/models"
)

func (app *App) MonitorInterrupts() {
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
			app.server.Shutdown(ctx)
		}
	}()
}

func (app *App) ConnectDb() error {
	initDb := !existDir(&app.databaseFile)

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

func (app *App) CloseDB() {
	log.Println("Closing database connection")
	if app.database != nil {
		app.database.Close()
	}
}

func (app *App) InitServer() {
	app.server = &http.Server{
		Addr:         app.addr,
		Handler:      app.Routes(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func (app *App) RunServer() {
	log.Printf("Listening on %s\n", app.addr)

	if err := app.server.ListenAndServeTLS(app.tlsCert, app.tlsKey); err != nil {
		if err != http.ErrServerClosed {
			log.Println("The error below raised after shutdown:")
			log.Println(err)
		}
	}
}
