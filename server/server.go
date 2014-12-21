package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers() {
	r := mux.NewRouter()
	for k, v := range routes {
		r.HandleFunc(k, v)
		http.Handle(k, r)
	}
	http.HandleFunc("/static/", StaticHandler)
}
