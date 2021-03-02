package subscriptions

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"subalogue/app/validators"
	"subalogue/db"
	"subalogue/helpers"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	query := db.GetQuery()

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
	err = decoder.Decode(&subscriptionParams)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	// Validator should accept the generalized Subscription struct
	// TODO: Shouldn't we check subscriptionParams or the request body?
	valid, paramErrors := validators.ValidateSubscription(
		db.Subscription{
			Name:        subscriptionParams.Name,
			Url:         subscriptionParams.Url,
			Price:       subscriptionParams.Price,
			Recurring:   subscriptionParams.Recurring,
			BillingDate: subscriptionParams.BillingDate,
		})

	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(paramErrors)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	subscriptionParams.UserID = user.ID
	createdSub, err := query.CreateSubscription(ctx, subscriptionParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdSub)
	if err != nil {
		log.Fatal(err)
	}
}
