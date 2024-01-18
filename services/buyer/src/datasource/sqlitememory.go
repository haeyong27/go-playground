package datasource

import (
	"context"
	"riad/ent"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteMemory struct{}

func NewSqliteMemory() *SqliteMemory {
	return &SqliteMemory{}
}

func (s *SqliteMemory) Connect() *ent.Client {
	client, err := ent.Open("sqlite3", "file::memory:?_fk=1")

	if err != nil {
		panic(err)
	}
	client.Schema.Create(context.Background())
	return client
}
