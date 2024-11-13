package main

import (
	"log"
	"net/http"
	errorHandler "workout-tracker/api/handlers/error"
	userHandler "workout-tracker/api/handlers/user"
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

	mux.HandleFunc("POST /users/{user_id}", userHandler.CreateUser)
	mux.HandleFunc("GET /users/{user_id}", userHandler.GetUser)
	mux.HandleFunc("PUT /users/{user_id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{user_id}", userHandler.DeleteUser)

	mux.HandleFunc("/", errorHandler.NotFound)

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(mux),
	}

	log.Printf("[INFO] Server starting on %s", s.addr)

	return server.ListenAndServe()
}
