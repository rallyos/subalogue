package subscriptions

import (
	"context"
	"net/http"
	"strconv"
	"subalogue/db"
	"subalogue/helpers"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := context.Background()
	query := db.GetQuery()
	var subscriptionParams db.DeleteSubscriptionParams

	w.Header().Set("Content-Type", "application/json")
	subscriptionID, _ := strconv.ParseInt(vars["id"], 10, 32)

	user, err := helpers.GetSessionUser(r)
	if err != nil {
		// May have some side effects but let us ignore actual errors for now
		// as all of them for now are due to invalid session anyway
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	subscriptionParams.UserID = user.ID
	subscriptionParams.ID = int32(subscriptionID)

	_, err = query.DeleteSubscription(ctx, subscriptionParams)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
