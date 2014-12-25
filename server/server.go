package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jd1123/johnnydiabetic.com/blog"
	"github.com/jd1123/johnnydiabetic.com/middleware"
)

func RegisterHandlers() {
	// Build a new router
	r := mux.NewRouter()

	// This is the root app
	for k, v := range routes {
		r.HandleFunc(k, v)
		http.Handle(k, middleware.LogRequest(r))
		log.Println("registering", k)
	}

	// Register Apps here
	blog.RegisterHandlers(r)

	// Static files
	http.HandleFunc("/static/", StaticHandler)
}
