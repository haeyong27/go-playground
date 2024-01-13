package query

import (
	"context"
	"riad/ent"
	"riad/src/datasource"
)

func GetUser() []*ent.User {
	db := datasource.NewDatabase("sqlite").Connect()
	ctx := context.Background()
	users, err := db.User.Query().All(ctx)
	if err != nil {
		panic(err)
	}
	return users
}
