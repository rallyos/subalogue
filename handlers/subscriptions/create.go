package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"
	"subalogue/db"
	"subalogue/session"
)

var ctx = context.Background()

func Create(w http.ResponseWriter, r *http.Request) {
	// TODO Validate r.Body
	// https://github.com/gorilla/schema If problems arise

	var subscription_params db.CreateUserSubscriptionParams

	username, err := session.Get(r, "username")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := db.Query.FindUserByUsername(ctx, username.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&subscription_params)

	subscription_params.UserID = user.ID
	db.Query.CreateUserSubscription(ctx, subscription_params)

	w.WriteHeader(http.StatusCreated)
}
