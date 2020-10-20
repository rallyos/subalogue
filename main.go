package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"github.com/coreos/go-oidc"
	// "github.com/dmralev/subalogue/app"
	"github.com/dmralev/subalogue/auth"
	"github.com/dmralev/subalogue/repositories"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"

	"github.com/gorilla/sessions"
	"net/http"
)

var query *db.Queries
var ctx = context.Background()
var Store = sessions.NewFilesystemStore("", []byte("auth-session"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Generate random state
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := Store.Get(r, "auth-session")
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

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		log.Printf("no token found: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: "4uNQLwCkcWPYKoTmSLvot6L2u39VOhWi",
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)

	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["username"] = profile["nickname"]
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/me/subscriptions", http.StatusSeeOther)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong\n"))
}

func CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {

	// DRY
	session, err := Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var username = session.Values["username"].(string)
	user, err := query.FindUserByUsername(ctx, username)

	var subscription_json db.CreateUserSubscriptionParams

	dec := json.NewDecoder(r.Body)
	dec.Decode(&subscription_json) // https://github.com/gorilla/schema If problems arise

	subscription_json.UserID = user.ID
	query.CreateUserSubscription(ctx, subscription_json)

	w.WriteHeader(http.StatusCreated)
}

func ListSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Check if null and flip the Valid key when true
	session, err := Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var username = session.Values["username"].(string)
	user, err := query.FindUserByUsername(ctx, username)

	subscriptions, err := query.ListUserSubscriptions(ctx, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(subscriptions)
}

func main() {

	conn, err := sql.Open("postgres", "user=subalogue password=subalogue dbname=subalogue_development host=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	query = db.New(conn)
	gob.Register(map[string]interface{}{})

	r := mux.NewRouter()

	r.HandleFunc("/ping", PingHandler)
	r.HandleFunc("/me/subscriptions", CreateSubscriptionHandler).Methods("POST")
	r.HandleFunc("/me/subscriptions", ListSubscriptionsHandler).Methods("GET")
	r.HandleFunc("/callback", CallbackHandler)
	r.HandleFunc("/login", LoginHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
