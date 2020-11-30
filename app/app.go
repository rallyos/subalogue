package app

import (
	"fmt"
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
	wd, _ := os.Getwd()
	fmt.Println(wd)
	chep
	envPath := fmt.Sprintf("%s/%s", wd, ".env."+os.Getenv("SUBALOGUE_ENV"))
	log.Println(envPath)
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) Run() {
	s.Router.Use(middlewares.LoggingMiddleware)
	s.Router.Use(middlewares.CORSMiddleware)
	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
