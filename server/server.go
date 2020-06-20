package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yudhasubki/go-skeleton/container"

	"github.com/yudhasubki/go-skeleton/config"
	"github.com/yudhasubki/go-skeleton/libs/api"
	"github.com/yudhasubki/go-skeleton/router"
)

type Server struct {
	DB     *sql.DB
	Config *config.Config
	Router *router.RouterHandler
}

func (s *Server) Shutdown(done chan os.Signal, svc container.ServiceRegistry) {
	<-done
	defer os.Exit(0)
	log.Println("signal shutdown received.")
	time.Sleep(time.Duration(5) * time.Second)
	log.Println("Clean up all resources...")
	svc.Shutdown()
	log.Println("Server shutdown properly...")
}

func (s *Server) EnableGracefulShutdown(svc container.ServiceRegistry) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go s.Shutdown(done, svc)
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
