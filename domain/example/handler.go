package example

import (
	"log"
	"net/http"
)

type Handler struct {
	Service *Service `inject:"service"`
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	log.Println("Server is alive.")
}
