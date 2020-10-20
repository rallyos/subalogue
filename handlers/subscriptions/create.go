package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"
	"subalogue/db"
	"subalogue/session"
)

var ctx = context.Background()

func CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {

	// DRY
	session, err := session.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var username = session.Values["username"].(string)
	user, err := db.Query.FindUserByUsername(ctx, username)

	var subscription_json db.CreateUserSubscriptionParams

	dec := json.NewDecoder(r.Body)
	dec.Decode(&subscription_json) // https://github.com/gorilla/schema If problems arise

	subscription_json.UserID = user.ID
	db.Query.CreateUserSubscription(ctx, subscription_json)

	w.WriteHeader(http.StatusCreated)
}
