package example

import (
	"database/sql"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository(Db *sql.DB) *Repository {
	return &Repository{
		Db,
	}
}
