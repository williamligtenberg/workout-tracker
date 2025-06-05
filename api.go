package main

import (
	"log"
	"net/http"

	errorHandler "github.com/williamligtenberg/workout-tracker/handlers/error"
	userHandler "github.com/williamligtenberg/workout-tracker/handlers/user"
	"github.com/williamligtenberg/workout-tracker/middleware"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		userHandler.CreateUser(w, r)
	})
	mux.HandleFunc("GET /users/{user_id}", userHandler.GetUser)
	mux.HandleFunc("DELETE /users/{user_id}", userHandler.DeleteUser)

	mux.HandleFunc("/", errorHandler.NotFound)

	middlewareChain := MiddlewareChain(
		middleware.RequestLoggerMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(mux),
	}

	log.Printf("[INFO] Server starting on %s", s.addr)

	return server.ListenAndServe()
}
