package server

import (
	"log"
	"net/http"
	"time"

	v1 "go-restful/internal/server/v1"

	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

//New server config
func New(port string) (*Server, error) {
	router := chi.NewRouter()

	// API routes version 1.
	router.Mount("/api/v1", v1.New())

	newServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: newServer}

	return &server, nil
}

//Close the server resources
func (server *Server) Close() error {
	return nil
}

//Start the server
func (server *Server) Start() {
	log.Printf("Server running on localhost%s", server.server.Addr)
	log.Fatal(server.server.ListenAndServe())
}
