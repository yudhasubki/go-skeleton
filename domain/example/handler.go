package example

import (
	"database/sql"
	"log"
	"net/http"

	conf "github.com/yudhasubki/go-skeleton/config"
)

type Handler struct {
	Service *Service
}

func NewHandler(Db *sql.DB, cfg *conf.Config) *Handler {
	return &Handler{
		Service: NewService(Db, cfg),
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	log.Println("Server is alive.")
}
