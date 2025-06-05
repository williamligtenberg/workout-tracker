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

	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users/{user_uuid}", userHandler.GetUser)
	mux.HandleFunc("DELETE /users/{user_uuid}", userHandler.DeleteUser)

	mux.HandleFunc("/", errorHandler.NotFound)

	handlerWithMiddleware := middleware.LoggingMiddleware(mux)

	server := http.Server{
		Addr:    s.addr,
		Handler: handlerWithMiddleware,
	}

	log.Printf("[INFO] Server starting on %s", s.addr)

	return server.ListenAndServe()
}
