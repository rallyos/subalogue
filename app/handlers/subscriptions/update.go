package subscriptions

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"subalogue/app/validators"
	"subalogue/db"
	"subalogue/helpers"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	query := db.GetQuery()

	vars := mux.Vars(r)
	subscriptionID, _ := strconv.ParseInt(vars["id"], 10, 32)

	var subscriptionParams db.UpdateSubscriptionParams

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
		log.Fatal(err)
	}

	valid, paramErrors := validators.ValidateSubscription(
		db.Subscription{
			Name:  subscriptionParams.Name,
			Url:   subscriptionParams.Url,
			Price: subscriptionParams.Price})

	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(paramErrors)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	subscriptionParams.UserID = user.ID
	subscriptionParams.ID = int32(subscriptionID)

	updatedSub, err := query.UpdateSubscription(ctx, subscriptionParams)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedSub)
	if err != nil {
		log.Fatal(err)
	}
}
