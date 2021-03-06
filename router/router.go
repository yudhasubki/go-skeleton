package router

import (
	"github.com/gorilla/mux"
	"github.com/yudhasubki/go-skeleton/domain/example"
)

type RouterHandler struct {
	Handler *example.Handler `inject:"handler"`
}

func Router(rh *RouterHandler) *mux.Router {
	router := mux.NewRouter()
	return router
}
