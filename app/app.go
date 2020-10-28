package app

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"database/sql"
	"subalogue/db"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func (s *Server) Initialize() {
	setEnv()
	s.Router = mux.NewRouter()
	db.Init()
	s.DB = db.GetConnection()
	s.routes()
}

func setEnv() {
	godotenv.Load(".env." + os.Getenv("SUBALOGUE_ENV"))
}

func (s *Server) Run() {
	// TODO addr param
	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
