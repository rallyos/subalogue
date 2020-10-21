package session

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.FilesystemStore
)

func init() {
	// https://github.com/auth0-samples/auth0-golang-web-app/tree/master/01-Login
	Store = sessions.NewFilesystemStore("", []byte("subalogue-auth-session"))
	gob.Register(map[string]interface{}{})
}

func Get(r *http.Request, key string) (interface{}, error) {
	session, err := Store.Get(r, "subalogue-auth-session")
	if err != nil {
		return nil, err
	}
	value := session.Values[key]
	return value, nil
}
