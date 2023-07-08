package models

import "database/sql"

type Post struct {
	ID        int
	Sentence  string
	CreatedAt string
	UpdatedAt string
	DeletedAt sql.NullString
}
