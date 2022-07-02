package middlewares

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func CORSMiddleware(next http.Handler) http.Handler {
	if os.Getenv("SUBALOGUE_ENV") == "development" {
		origins := handlers.AllowedOrigins([]string{"http://localhost:8000"})
		methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions})
		headers := handlers.AllowedHeaders([]string{"Content-Type", "Cookie"})
		creds := handlers.AllowCredentials()
		return handlers.CORS(origins, methods, headers, creds)(next)
	}
	return next
}
