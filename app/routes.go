package app

import (
	"subalogue/app/handlers"
	"subalogue/app/handlers/auth"
	"subalogue/app/handlers/subscriptions"
)

func (s *Server) routes() {
	s.Router.HandleFunc("/ping", handlers.PingHandler)
	s.Router.HandleFunc("/api/v1/me/subscriptions", subscriptions.Create).Methods("POST", "OPTIONS")
	s.Router.HandleFunc("/api/v1/me/subscriptions", subscriptions.List).Methods("GET")
	s.Router.HandleFunc("/api/v1/me/subscriptions/{id:[0-9]+}", subscriptions.Update).Methods("PUT")
	s.Router.HandleFunc("/auth/callback", auth.Callback)
	s.Router.HandleFunc("/auth/login", auth.Login)
}
