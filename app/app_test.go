package app

import (
	"bytes"
	"fmt"
	"net/http"
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
	tables := []string{"users", "subscriptions"}
	for i := 0; i < len(tables); i++ {
		server.DB.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", tables[i]))
		server.DB.Exec(fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART", tables[i]))
		server.DB.Exec("UPDATE %s SET id = DEFAULT", tables[i])
	}
}

func populate() {
	userRow := server.DB.QueryRow(`INSERT INTO users(username) VALUES ('dmralev') RETURNING username`)
	userRow.Scan(&username)
	subscriptionRow := server.DB.QueryRow(`INSERT INTO subscriptions(name, url, price, username) VALUES ('Brilliant', 'https://brilliant.org', 5900, 'dmralev') RETURNING *`)
	subscriptionRow.Scan(&username)
}

func setSessionKey(k, v string, req *http.Request, t *testing.T) {
	s, err := session.Store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Error(err)
	}
	s.Values[k] = v
}

func TestPing(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	server.Router.ServeHTTP(w, req)

	expectedCode, expected_body := 200, "Pong"
	statusCode, body := w.Result().StatusCode, w.Body.String()

	is.Equal(statusCode, expectedCode)
	is.Equal(body, expected_body)
}

func TestListSubscriptions(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest("GET", "/api/v1/me/subscriptions", nil)
	setSessionKey("username", username, req, t)

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	expectedCode := 200
	statusCode := w.Result().StatusCode

	is.Equal(statusCode, expectedCode)
}

func TestCreateSubscriptions(t *testing.T) {
	is := is.New(t)

	jsonStr := []byte(`{"name": "Brain.fm", "url": "https://brain.fm", "price": 299}`)
	req := httptest.NewRequest("POST", "/api/v1/me/subscriptions", bytes.NewBuffer(jsonStr))
	setSessionKey("username", username, req, t)

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	expectedCode := 201
	statusCode := w.Result().StatusCode

	is.Equal(statusCode, expectedCode)
}

func TestUpdateSubscriptions(t *testing.T) {
	is := is.New(t)

	jsonStr := []byte(`{"name": "Brain.fm", "url": "https://brain.fm", "price": 399}`)
	req := httptest.NewRequest("PUT", "/api/v1/me/subscriptions/1", bytes.NewBuffer(jsonStr))
	setSessionKey("username", username, req, t)

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	expectedCode := 200
	statusCode := w.Result().StatusCode

	is.Equal(statusCode, expectedCode)
}

func TestDeleteSubscriptions(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest("DELETE", "/api/v1/me/subscriptions/1", nil)
	setSessionKey("username", username, req, t)

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	expectedCode := 204
	statusCode := w.Result().StatusCode

	is.Equal(statusCode, expectedCode)
}
