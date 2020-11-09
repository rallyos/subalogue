package handlers

import (
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Pong"))
	if err != nil {
		log.Fatal(err)
	}
}
