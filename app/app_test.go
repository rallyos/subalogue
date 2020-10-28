package app

import (
	"bytes"
	"net/http/httptest"
	"os"
	"subalogue/session"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/matryer/is"
)

var server *Server
var username string

func init() {
	godotenv.Load("../.env.test")
	server = &Server{}
	server.Initialize()
	preSetup()
}

func preSetup() {
	migrateDB()
	clearDB()
	populate()
}

func migrateDB() {
	t := testing.T{}
	m, err := migrate.New("file://../db/migrations", os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Log(err)
	}
	m.Up()
}

func clearDB() {
	server.DB.Exec("TRUNCATE users CASCADE")
	server.DB.Exec("ALTER SEQUENCE users_id_seq RESTART")
	server.DB.Exec("UPDATE users SET id = DEFAULT")
}

func populate() {
	userRow := server.DB.QueryRow(`INSERT INTO users(username) VALUES ('dmralev') RETURNING username`)
	userRow.Scan(&username)
}

func TestPing(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	server.Router.ServeHTTP(w, req)

	expected_code, expected_body := 200, "Pong"
	status_code, body := w.Result().StatusCode, w.Body.String()

	is.Equal(status_code, expected_code)
	is.Equal(body, expected_body)
}

func TestCreateSubscriptions(t *testing.T) {
	is := is.New(t)

	jsonStr := []byte(`{"name": "HBOGO", "price": 799}`)
	req := httptest.NewRequest("POST", "/api/v1/me/subscriptions", bytes.NewBuffer(jsonStr))

	s, err := session.Store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Error(err)
	}
	s.Values["username"] = username

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	expected_code := 201
	status_code := w.Result().StatusCode

	is.Equal(status_code, expected_code)
}

func TestListSubscriptions(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest("GET", "/api/v1/me/subscriptions", nil)

	s, err := session.Store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Error(err)
	}
	s.Values["username"] = username

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	expected_code := 200
	status_code := w.Result().StatusCode

	is.Equal(status_code, expected_code)
}
