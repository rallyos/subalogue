package subscriptions

import (
	"context"
	"encoding/json"
	"net/http"
	"subalogue/db"
	"subalogue/session"
)

func List(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	query := db.GetQuery()

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

	subscriptions, err := query.ListUserSubscriptions(ctx, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(subscriptions)
}
