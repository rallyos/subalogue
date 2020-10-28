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

	// TODO Catch https://github.com/gorilla/sessions/issues/209#issuecomment-694341696
	username, err := session.Get(r, "username")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if username == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := query.FindUserByUsername(ctx, username.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	subscriptions, err := query.ListSubscriptions(ctx, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(subscriptions) == 0 {
		subscriptions = make([]db.Subscription, 0)
	}

	subs := map[string][]db.Subscription{
		"subscriptions": subscriptions,
	}

	json.NewEncoder(w).Encode(subs)
}
