package router

import (
	"github.com/gorilla/mux"
	"github.com/yudhasubki/go-skeleton/domain/example"
)

type RouterHandler struct {
	Handler *example.Handler
}

func Router(rh *RouterHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", rh.Handler.Health).Methods("GET")

	return router
}
