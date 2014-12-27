package server

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jd1123/johnnydiabetic.com/blog"
	"github.com/jd1123/johnnydiabetic.com/config"
	"github.com/jd1123/johnnydiabetic.com/middleware"
)

var S sessions.Store

func init() {
	S = sessions.NewCookieStore(config.AuthenticationKey(), config.EncryptionKey())
	gob.Register(&User{})
	/*
		S.Options = &sessions.Options{
			Domain:   "localhost",
			Path:     "/",
			MaxAge:   3600 * 8,
			HttpOnly: true,
		}
	*/
}

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
