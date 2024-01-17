package query

import (
	"context"
	"riad/ent"
	"riad/services/buyer/src/datasource"
)

func GetUser() []*ent.User {
	ctx := context.Background()
	users, err := datasource.Client.User.Query().All(ctx)
	if err != nil {
		panic(err)
	}
	return users
}

func CreateUser() *ent.User {
	ctx := context.Background()
	user, err := datasource.Client.User.Create().SetAge(30).Save(ctx)
	if err != nil {
		panic(err)
	}
	return user
}
