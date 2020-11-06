package app

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"database/sql"
	"subalogue/db"

	"github.com/gorilla/handlers"
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
	godotenv.Load(".env." + os.Getenv("SUBALOGUE_ENV"))
}

func (s *Server) Run() {
	// TODO addr param
	// log.Fatal(http.ListenAndServe(":8000", s.Router))
	origins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Cookie"})
	creds := handlers.AllowCredentials()
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(origins, methods, headers, creds)(s.Router)))
}
