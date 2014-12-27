package server

import "net/http"

type HandlerFunction func(w http.ResponseWriter, r *http.Request)

var routes = map[string]HandlerFunction{
	"/":           IndexHandler,
	"/about/":     AboutHandler,
	"/login/":     LoginHandler,
	"/logout/":    LogoutHandler,
	"/robots.txt": RobotsHandler,
}
