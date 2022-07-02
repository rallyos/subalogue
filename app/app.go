package app

import (
	"log"
	"net/http"
	"os"

	"database/sql"
	"subalogue/app/middlewares"
	"subalogue/db"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

// Needs rework soon
func setEnv() {
	if os.Getenv("SUBALOGUE_ENV") == "production" {
		err := godotenv.Load(".env.production")
		if err != nil {
			log.Fatal(err.Error())
		}

	}
}

func (s *Server) Run() {
	s.Router.Use(middlewares.LoggingMiddleware)
	s.Router.Use(middlewares.CORSMiddleware)
	log.Fatal(http.ListenAndServe(":3000", s.Router))
}
