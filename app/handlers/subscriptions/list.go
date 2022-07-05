package subscriptions

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/shiftingphotons/subalogue/db"
	"github.com/shiftingphotons/subalogue/helpers"
)

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	query := db.GetQuery()

	user, err := helpers.GetSessionUser(r)
	if err != nil {
		// May have some side effects but let us ignore actual errors for now
		// as all of them for now are due to invalid session anyway
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	subscriptions, err := query.ListSubscriptions(ctx, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	subs := map[string][]db.Subscription{
		"subscriptions": subscriptions,
	}

	err = json.NewEncoder(w).Encode(subs)
	if err != nil {
		log.Fatal(err)
	}
}
