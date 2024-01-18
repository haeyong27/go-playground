package query

import (
	"buyer/src/datasource"
	"context"
	"riad/ent"
)

func GetUser() []*ent.User {
	ctx := context.Background()
	users, err := datasource.Client.User.Query().All(ctx)
	if err != nil {
		panic(err)
	}
	return users
}

func CreateUser(name string, age int) *ent.User {
	ctx := context.Background()
	user, err := datasource.Client.User.Create().SetName(name).SetAge(age).Save(ctx)
	if err != nil {
		panic(err)
	}
	return user
}
