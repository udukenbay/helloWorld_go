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

// Запуск сервера
func (s *Server) Start(router *mux.Router) error {
	// Установка префикса для статических файлов
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("views/"))))
	http.Handle("/", router)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}

// Остановка сервера
func (s *Server) Stop() error {
	return nil
}
