package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[INFO] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("[INFO] Completed in %v", time.Since(start))
	})
}
