package app

import (
	"net/http/httptest"
	"testing"
)

var server *Server

func Init() {
	server = &Server{}
	server.Initialize()
}

func TestPing(t *testing.T) {
	Init()
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	server.Router.ServeHTTP(w, req)

	expected_code, expected_body := 200, "Pong"
	status_code, body := w.Result().StatusCode, w.Body.String()

	if status_code != expected_code {
		t.Errorf("Expected %d, got %d", expected_code, status_code)
	}

	if body != expected_body {
		t.Errorf("Expected %s, got %s", expected_body, body)
	}
}
