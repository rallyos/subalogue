package subscriptions

import (
	"encoding/json"
	"net/http"
	"subalogue/db"
	"subalogue/session"
)

func List(w http.ResponseWriter, r *http.Request) {
	username, err := session.Get(r, "username")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := db.Query.FindUserByUsername(ctx, username.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	subscriptions, err := db.Query.ListUserSubscriptions(ctx, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(subscriptions)
}
