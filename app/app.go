package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

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
	s.DB = db.Init()
	s.routes()
}

func setEnv() {
	env := os.Getenv("SUBALOGUE_ENV")
	godotenv.Load(".env." + env)
	fmt.Println(".env." + env + " Loaded.")
}

func (s *Server) Run() {
	// TODO addr param
	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
