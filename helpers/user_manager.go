package helpers

import (
	"context"
	"subalogue/db"
)

func CreateUser(username string) {
	query := db.GetQuery()
	ctx := context.Background()
	userFound, _ := query.FindUserByUsername(ctx, username)
	if userFound.ID == 0 {
		query.CreateUser(ctx, username)
	}
}
