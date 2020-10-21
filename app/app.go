package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func (s *Server) Initialize() {
	s.router = mux.NewRouter()
	s.routes()
}

func (s *Server) Run() {
	// TODO addr param
	log.Fatal(http.ListenAndServe(":8000", s.router))
}
