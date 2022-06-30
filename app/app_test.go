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

const sessKey string = "username"
const firstSubURL string = "/api/v1/me/subscriptions/1"
const subListURL string = "/api/v1/me/subscriptions"

func init() {
	t := testing.T{}

	err := godotenv.Load("/app/.env.test")
	if err != nil {
		t.Log(err)
	}

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

	m, err := migrate.New("file:///app/db/migrations", os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Log(err)
	}

	err = m.Up()
	if err != nil {
		t.Log(err)
	}
}

func clearDB() {
	t := testing.T{}

	tables := []string{"users", "subscriptions"}
	for i := 0; i < len(tables); i++ {
		_, err := server.DB.Exec(fmt.Sprintf("TRUNCATE %[1]s CASCADE; ALTER SEQUENCE %[1]s_id_seq RESTART; UPDATE %[1]s SET id = DEFAULT", tables[i]))
		if err != nil {
			t.Log(err)
		}
	}
}

const createUserQuery string = `INSERT INTO users(username) VALUES ('dmralev') RETURNING username`
const createSubQuery string = `INSERT INTO subscriptions(name, url, price, username) VALUES ('Brilliant', 'https://brilliant.org', 5900, 'dmralev') RETURNING *`

func populate() {
	t := testing.T{}

	userRow := server.DB.QueryRow(createUserQuery)
	if err := userRow.Scan(&username); err != nil {
		t.Log(err)
	}
	subscriptionRow := server.DB.QueryRow(createSubQuery)
	if err := subscriptionRow.Scan(&username); err != nil {
		t.Log(err)
	}
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

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	server.Router.ServeHTTP(w, req)

	expectedCode, expected_body := http.StatusOK, "Pong"
	statusCode, body := w.Result().StatusCode, w.Body.String()

	is.Equal(statusCode, expectedCode)
	is.Equal(body, expected_body)
}

func TestListSubscriptions(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest(http.MethodGet, subListURL, nil)
	setSessionKey(sessKey, username, req, t)

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	statusCode := w.Result().StatusCode

	is.Equal(statusCode, http.StatusOK)
}

func TestCreateSubscriptions(t *testing.T) {
	is := is.New(t)

	jsonStr := []byte(`{"name": "Brain.fm", "url": "https://brain.fm", "price": 299, "billing_date": "2021-03-01T00:00:00Z", "recurring": "monthly"}`)
	req := httptest.NewRequest(http.MethodPost, subListURL, bytes.NewBuffer(jsonStr))
	setSessionKey(sessKey, username, req, t)

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	statusCode := w.Result().StatusCode

	is.Equal(statusCode, http.StatusCreated)
}

func TestUpdateSubscriptions(t *testing.T) {
	is := is.New(t)

	jsonStr := []byte(`{"name": "Brain.fm", "url": "https://brain.fm", "price": 399, "billing_date": "2021-03-01T00:00:00Z", "recurring": "yearly"}`)
	req := httptest.NewRequest(http.MethodPut, firstSubURL, bytes.NewBuffer(jsonStr))
	setSessionKey(sessKey, username, req, t)

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	statusCode := w.Result().StatusCode

	is.Equal(statusCode, http.StatusOK)
}

func TestDeleteSubscriptions(t *testing.T) {
	is := is.New(t)

	req := httptest.NewRequest(http.MethodDelete, firstSubURL, nil)
	setSessionKey(sessKey, username, req, t)

	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	statusCode := w.Result().StatusCode

	is.Equal(statusCode, http.StatusNoContent)
}
