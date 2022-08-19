package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Port int
}

func New(port int) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Start(router *mux.Router) error {
	http.Handle("/", router)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}

func (s *Server) Stop() error {
	return nil
}
