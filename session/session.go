package session

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var (
	Store *sessions.FilesystemStore
)

func init() {
	// https://github.com/auth0-samples/auth0-golang-web-app/tree/master/01-Login
	Store = sessions.NewFilesystemStore("", []byte(os.Getenv("SESSION_KEY")))
	gob.Register(map[string]interface{}{})
}

func Get(r *http.Request, key string) (interface{}, error) {
	session, err := Store.Get(r, os.Getenv("SESSION_KEY"))
	if err != nil {
		return nil, err
	}
	value := session.Values[key]
	return value, nil
}
