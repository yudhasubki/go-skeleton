package server

import (
	"database/sql"
	"net/http"

	"github.com/yudhasubki/go-skeleton/config"
	"github.com/yudhasubki/go-skeleton/libs/api"
	"github.com/yudhasubki/go-skeleton/router"
)

type Server struct {
	DB     *sql.DB
	Config *config.Config
	Router *router.RouterHandler
}

func (s *Server) Serve() {
	routes := router.Router(s.Router)
	routes.HandleFunc("/health", s.HealthCheck).Methods("GET")
	http.ListenAndServe(":"+s.Config.ServerPort, routes)
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := s.DB.Ping(); err != nil {
		api.Write(w, api.Response("500", "server is unhealthy", nil))
		return
	}
	api.Write(w, api.Response("200", "server is alive", nil))
}
