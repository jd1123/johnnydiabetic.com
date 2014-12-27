package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jd1123/johnnydiabetic.com/blog"
	"github.com/jd1123/johnnydiabetic.com/config"
	"github.com/jd1123/johnnydiabetic.com/middleware"
)

var Store = sessions.NewCookieStore(config.AuthenticationKey(), config.EncryptionKey())

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

	// 404 Handler
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// Static files
	http.HandleFunc("/static/", StaticHandler)
}
