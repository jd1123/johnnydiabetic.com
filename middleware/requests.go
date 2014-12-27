package middleware

import (
	"log"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("URI Requested:", r.RequestURI, "User Agent:", r.Header["User-Agent"], "Remote Address:", r.Header["X-Real-IP"])
		next.ServeHTTP(w, r)
	})
}
