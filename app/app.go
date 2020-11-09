package app

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

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
	setEnv()
	db.Init()
	s.Router = mux.NewRouter()
	s.DB = db.GetConnection()
	s.routes()
}

func setEnv() {
	err := godotenv.Load(".env." + os.Getenv("SUBALOGUE_ENV"))
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Run() {
	s.Router.Use(middlewares.LoggingMiddleware)
	s.Router.Use(middlewares.CORSMiddleware)
	// TODO addr param
	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
