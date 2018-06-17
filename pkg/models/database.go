package models

import (
	"database/sql"
	"time"
)

type Database struct {
	*sql.DB
}

func (db *Database) GetSnippet(id int) (*Snippet, error) {
	if id == 123 {
		snippet := &Snippet{
			ID:      id,
			Title:   "Example Title",
			Content: "Example Content",
			Created: time.Now(),
			Expires: time.Now(),
		}
		return snippet, nil
	}

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
