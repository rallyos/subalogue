package app

import (
	"net/http"
	"subalogue/app/handlers"
	"subalogue/app/handlers/auth"
	"subalogue/app/handlers/subscriptions"
)

func (s *Server) routes() {
	s.Router.HandleFunc("/ping", handlers.PingHandler)
	s.Router.HandleFunc("/api/v1/me/subscriptions", subscriptions.Create).Methods(http.MethodPost, http.MethodOptions)
	s.Router.HandleFunc("/api/v1/me/subscriptions", subscriptions.List).Methods(http.MethodGet)
	s.Router.HandleFunc("/api/v1/me/subscriptions/{id:[0-9]+}", subscriptions.Update).Methods(http.MethodPut, http.MethodOptions)
	s.Router.HandleFunc("/api/v1/me/subscriptions/{id:[0-9]+}", subscriptions.Delete).Methods(http.MethodDelete, http.MethodOptions)
	s.Router.HandleFunc("/auth/callback", auth.Callback)
	s.Router.HandleFunc("/auth/login", auth.Login)
}
