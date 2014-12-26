package middleware

import (
	"log"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("User Agent:", r.Header["User-Agent"], "Remote Address:", r.Header["X-Real-Ip"])
		next.ServeHTTP(w, r)
	})
}
