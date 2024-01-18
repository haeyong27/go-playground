package datasource

import "riad/ent"

var Client *ent.Client

type IDatabase interface {
	Connect() *ent.Client
}

type DataSource struct {
	client *ent.Client
	dbType string
}

func NewDatabase(dbType string) *DataSource {
	ds := new(DataSource)
	ds.dbType = dbType
	return ds
}

func (ds *DataSource) Connect() {
	var client IDatabase
	switch ds.dbType {
	case "s":
		client = NewSqlite()
	case "m":
		client = NewSqliteMemory()
	default:
		client = NewSqliteMemory()
	}

	Client = client.Connect()
}
