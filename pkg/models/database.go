package models

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Database struct {
	*sql.DB
}

var ErrDuplicateEmail = errors.New("models: email address already in use")

func (db *Database) GetSnippet(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE id = ? AND expires > datetime('now')`

	row := db.QueryRow(stmt, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

func (db *Database) LatestSnippets() (Snippets, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > datetime('now') ORDER BY created DESC LIMIT 10`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := Snippets{}

	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (db *Database) InsertSnippet(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES (?, ?, datetime('now'), datetime('now', ?))`

	result, err := db.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (db *Database) InitializeDb() error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt := `
	CREATE TABLE snippets (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created DATETIME NOT NULL,
		expires DATETIME NOT NULL
	);
	  
	CREATE INDEX idx_snippets_created ON snippets(created);

	CREATE TABLE users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password CHAR(60) NOT NULL,
		created DATETIME NOT NULL
	);
	`

	_, err = tx.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `INSERT INTO snippets (title, content, created, expires)
	VALUES (?, ?, datetime('now'), datetime('now', ?))`

	hokku := map[int][]interface{}{
		0: {"An old silent pond", "An old silent pond...\nA frog jumps into the pond.\nSplash! Silence again.\n", "+14 days"},
		1: {"Over the wintry forrest", "Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n", "+14 days"},
		2: {"First autumn morning", "First autumn morning\nthe mirror I stare into\nshows my father''s face.\n", "+14 days"},
	}

	for _, v := range hokku {
		_, err = tx.Exec(stmt, v...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) InsertUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, password, created)
	VALUES (?, ?, ?, datetime('now'))`

	_, err = db.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		dbErr := err.(sqlite3.Error)
		if dbErr.Code == sqlite3.ErrConstraint && dbErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return ErrDuplicateEmail
		}
	}
	return err
}
