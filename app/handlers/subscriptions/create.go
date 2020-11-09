package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"
	"subalogue/app/validators"
	"subalogue/db"
	"subalogue/helpers"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(http.StatusOK)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	query := db.GetQuery()

	// TODO Validate r.Body
	// https://github.com/gorilla/schema If problems arise
	var subscriptionParams db.CreateSubscriptionParams

	user, err := helpers.GetSessionUser(r)
	if err != nil {
		// May have some side effects but let us ignore actual errors for now
		// as all of them for now are due to invalid session anyway
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&subscriptionParams)

	// Validator should accept the generalized Subscription struct
	valid, paramErrors := validators.ValidateSubscription(
		db.Subscription{
			Name:  subscriptionParams.Name,
			Url:   subscriptionParams.Url,
			Price: subscriptionParams.Price})

	if valid == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(paramErrors)
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
