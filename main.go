package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dmralev/subalogue/repositories"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var query *db.Queries
var ctx = context.Background()

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong\n"))
}

func CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var subscription_json db.CreateUserSubscriptionsParams
	harcoded_user_id := sql.NullInt32{Int32: 1, Valid: true}

	dec := json.NewDecoder(r.Body)
	dec.Decode(&subscription_json) // https://github.com/gorilla/schema If problems arise

	subscription_json.UserID = harcoded_user_id
	query.CreateUserSubscriptions(ctx, subscription_json)

	w.WriteHeader(http.StatusCreated)
}

func ListSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Check if null and flip the Valid key when true
	harcoded_user_id := sql.NullInt32{Int32: 1, Valid: true}
	subscriptions, err := query.ListUserSubscriptions(ctx, harcoded_user_id)
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

	r := mux.NewRouter()

	r.HandleFunc("/ping", PingHandler)
	r.HandleFunc("/me/subscriptions", CreateSubscriptionHandler).Methods("POST")
	r.HandleFunc("/me/subscriptions", ListSubscriptionsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
