package app

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Initialize() {
	s.Router = mux.NewRouter()
	s.routes()
}

func (s *Server) Run() {
	// TODO addr param
	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
