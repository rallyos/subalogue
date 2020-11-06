package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"
	"subalogue/db"
	"subalogue/session"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
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

	if subscription_params.Url == "" {
		err_map := map[string]string{
			"name": "Name should not be empty.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err_map)
		return
	}

	subscription_params.UserID = user.ID
	created_sub, err := query.CreateSubscription(ctx, subscription_params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created_sub)
}
