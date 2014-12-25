package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jd1123/johnnydiabetic.com/blog"
)

func RegisterHandlers() {
	// Build a new router
	r := mux.NewRouter()

	// This is the root app
	for k, v := range routes {
		r.HandleFunc(k, v)
		http.Handle(k, r)
	}

	// Register Apps here
	blog.RegisterHandlers(r)

	// Static files
	http.HandleFunc("/static/", StaticHandler)
}
