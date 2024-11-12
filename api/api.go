package main

import (
	"log"
	"net/http"
	handlers "workout-tracker/api/handlers/user"
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
	router := http.NewServeMux()

	router.HandleFunc("POST /users", handlers.CreateUser)
	router.HandleFunc("GET /users/{user_id}", handlers.GetUser)
	router.HandleFunc("PUT /users/{user_id}", handlers.UpdateUser)
	router.HandleFunc("DELETE /users/{user_id}", handlers.DeleteUser)

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequireAuthMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	log.Printf("Server is now running on port %s", s.addr)

	return server.ListenAndServe()
}
