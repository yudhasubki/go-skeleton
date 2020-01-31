package example

import (
	"database/sql"

	conf "github.com/yudhasubki/go-skeleton/config"
)

type Service struct {
	Repo   *Repository
	Config *conf.Config
}

func NewService(Db *sql.DB, cfg *conf.Config) *Service {
	return &Service{
		Repo: NewRepository(Db),
	}
}
