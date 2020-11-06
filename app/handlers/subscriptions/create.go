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
	var subscriptionParams db.CreateSubscriptionParams

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
	decoder.Decode(&subscriptionParams)

	if subscriptionParams.Name == "" {
		errMap := map[string]string{
			"name": "Name should not be empty.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMap)
		return
	}

	if subscriptionParams.Url == "" {
		errMap := map[string]string{
			"url": "Url should not be empty.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errMap)
		return
	}

	subscriptionParams.UserID = user.ID

	createdSub, err := query.CreateSubscription(ctx, subscriptionParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdSub)
}
