package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"subalogue/helpers"
	"subalogue/session"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// https://github.com/auth0-samples/auth0-golang-web-app/tree/master/01-Login

	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := session.Store.Get(r, "subalogue-auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authenticator, err := helpers.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}