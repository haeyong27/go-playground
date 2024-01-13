package datasource

import "riad/ent"

var Client *ent.Client

type IDatabase interface {
	Connect() *ent.Client
}

func NewDatabase(dbType string) IDatabase {
	switch dbType {
	case "sqlite":
		return &Sqlite{}
	default:
		return &Sqlite{}
	}
}






