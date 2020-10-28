package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"
	"subalogue/db"
	"subalogue/session"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	query := db.GetQuery()

	// TODO Validate r.Body
	// https://github.com/gorilla/schema If problems arise
	var subscription_params db.CreateSubscriptionParams

	username, err := session.Get(r, "username")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := query.FindUserByUsername(ctx, username.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&subscription_params)

	if subscription_params.Name == "" {
		err_map := map[string]string{
			"name": "Name should not be empty.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err_map)
		return
	}

	subscription_params.UserID = user.ID
	query.CreateSubscription(ctx, subscription_params)

	w.WriteHeader(http.StatusCreated)
}
