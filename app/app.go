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
	// Hardcode root folder because there are issues with running the tests (expected behaviour in Go because of the working directory value when tests in other packages are run)
	// We use containers so we have a guarantee that this path will always be correct
	err := godotenv.Load("/app/.env." + os.Getenv("SUBALOGUE_ENV"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) Run() {
	s.Router.Use(middlewares.LoggingMiddleware)
	s.Router.Use(middlewares.CORSMiddleware)
	log.Fatal(http.ListenAndServe(":3000", s.Router))
}
