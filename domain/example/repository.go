package example

import (
	"database/sql"
)

type Repository struct {
	Db *sql.DB `inject:"database"`
}
