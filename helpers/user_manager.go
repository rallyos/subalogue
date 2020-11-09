package helpers

import (
	"context"
	"log"
	"net/http"
	"subalogue/db"
	"subalogue/session"
)

func CreateUser(username string) {
	query := db.GetQuery()
	ctx := context.Background()
	userFound, _ := query.FindUserByUsername(ctx, username)
	if userFound.ID == 0 {
		_, err := query.CreateUser(ctx, username)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetSessionUser(r *http.Request) (db.User, error) {
	query := db.GetQuery()
	ctx := context.Background()

	username, err := session.Get(r, "username")
	if err != nil || username == nil {
		return db.User{}, err
	}

	user, err := query.FindUserByUsername(ctx, username.(string))
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
