package main

import (
	"net/http"
)

// Server is a struct of the server
type Server struct {
	port string
}

// NewServer is the function to create the server
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

// Listen is a function to start the http ListenAndServe
func (s *Server) Listen() error {
	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		return err
	}
	return nil
}
