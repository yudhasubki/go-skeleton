package server

import (
	"net/http"

	"github.com/yudhasubki/go-skeleton/config"
	"github.com/yudhasubki/go-skeleton/libs/api"
	"github.com/yudhasubki/go-skeleton/router"
)

type Server struct {
	Config *config.Config
	Router *router.RouterHandler
}

func (s *Server) Serve() {
	routes := router.Router(s.Router)
	routes.HandleFunc("/health", s.HealthCheck).Methods("GET")
	http.ListenAndServe(":"+s.Config.ServerPort, routes)
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	api.Write(w, api.Response("200", "server is alive", nil))
}
