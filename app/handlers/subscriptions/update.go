package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"subalogue/db"
	"subalogue/session"

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
		err_map := map[string]string{
			"name": "Name should not be empty.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err_map)
		return
	}

	if subscriptionParams.Url == "" {
		err_map := map[string]string{
			"name": "Name should not be empty.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err_map)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	subscriptionParams.UserID = user.ID
	subscriptionParams.ID = int32(subscriptionID)

	_, err = query.UpdateSubscription(ctx, subscriptionParams)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
