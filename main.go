package main

import (
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"subalogue/db"
	"subalogue/handlers"
	"subalogue/handlers/auth"
	"subalogue/handlers/subscriptions"
	"subalogue/session"
)

func main() {
	session.Init()
	db.Init()

	r := mux.NewRouter()

	r.HandleFunc("/ping", handlers.PingHandler)
	// TODO subscriptions.Create
	r.HandleFunc("/me/subscriptions", subscriptions.CreateSubscriptionHandler).Methods("POST")
	r.HandleFunc("/me/subscriptions", subscriptions.ListSubscriptionsHandler).Methods("GET")
	r.HandleFunc("/auth/callback", auth.CallbackHandler)
	r.HandleFunc("/auth/login", auth.LoginHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
