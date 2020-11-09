package subscriptions

import (
	"context"
	"encoding/json"
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
	subscriptionID, err := strconv.ParseInt(vars["id"], 10, 32)

	//TODO: Validate
	var subscriptionParams db.UpdateSubscriptionParams

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
	subscriptionParams.ID = int32(subscriptionID)

	updatedSub, err := query.UpdateSubscription(ctx, subscriptionParams)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedSub)
}
