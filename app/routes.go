package app

import (
	"subalogue/handlers"
	"subalogue/handlers/auth"
	"subalogue/handlers/subscriptions"
)

func (s *Server) routes() {
	s.router.HandleFunc("/ping", handlers.PingHandler)
	s.router.HandleFunc("/me/subscriptions", subscriptions.Create).Methods("POST")
	s.router.HandleFunc("/me/subscriptions", subscriptions.List).Methods("GET")
	s.router.HandleFunc("/auth/callback", auth.Callback)
	s.router.HandleFunc("/auth/login", auth.Login)
}
