package middlewares

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func CORSMiddleware(next http.Handler) http.Handler {
	origins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Cookie"})
	creds := handlers.AllowCredentials()
	return handlers.CORS(origins, methods, headers, creds)(next)

}
