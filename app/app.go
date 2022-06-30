package app

import (
	"log"
	"net/http"

	"database/sql"
	"subalogue/app/middlewares"
	"subalogue/db"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func (s *Server) Initialize() {
	db.Init()
	s.Router = mux.NewRouter()
	s.DB = db.GetConnection()
	s.routes()
}

func (s *Server) Run() {
	s.Router.Use(middlewares.LoggingMiddleware)
	s.Router.Use(middlewares.CORSMiddleware)
	log.Fatal(http.ListenAndServe(":3000", s.Router))
}
