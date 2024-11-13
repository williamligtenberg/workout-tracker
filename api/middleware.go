package main

import (
	"encoding/json"
	"log"
	"net/http"
	handlers "workout-tracker/api/handlers/error"
)

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] Incoming request | Method: %s | Path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Token" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			log.Printf("[WARN] Unauthorized access attempt | Method: %s | Path: %s", r.Method, r.URL.Path)

			response := handlers.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Error:   "unauthorized",
				Message: "Authorization token is missing or invalid.",
			}

			if err := json.NewEncoder(w).Encode(response); err != nil {
				log.Printf("[ERROR] Failed to encode unauthorized response: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		log.Printf("[INFO] Authorized request | Method: %s | Path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}
