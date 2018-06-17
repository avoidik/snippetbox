package models

import (
	"time"
)

type Database struct{}

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
	return nil, nil
}
