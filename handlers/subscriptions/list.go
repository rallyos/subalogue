package subscriptions

import (
	"encoding/json"
	"log"
	"net/http"
	"subalogue/db"
	"subalogue/session"
)

func ListSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Check if null and flip the Valid key when true
	session, err := session.Store.Get(r, "subalogue-auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var username = session.Values["username"].(string)
	log.Println(username)
	user, err := db.Query.FindUserByUsername(ctx, username)

	subscriptions, err := db.Query.ListUserSubscriptions(ctx, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(subscriptions)
}
