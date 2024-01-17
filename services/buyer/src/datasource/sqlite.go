package datasource

import (
	"context"
	"riad/ent"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct{}

func NewSqlite() *Sqlite {
	return &Sqlite{}
}

func (s *Sqlite) Connect() *ent.Client {
	client, err := ent.Open("sqlite3", "file:riad.db?_fk=1")
	if err != nil {
		panic(err)
	}
	client.Schema.Create(context.Background())
	return client
}
