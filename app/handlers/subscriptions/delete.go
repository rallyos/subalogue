package subscriptions

import (
	"context"
	"net/http"
	"strconv"
	"subalogue/db"
	"subalogue/session"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	query := db.GetQuery()

	vars := mux.Vars(r)
	subscriptionID, err := strconv.ParseInt(vars["id"], 10, 32)

	//TODO: Validate
	var subscriptionParams db.DeleteSubscriptionParams

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

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
