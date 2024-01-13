package main

import (
	"riad/src/datasource"
)

func main() {

	database := datasource.NewDatabase("sqlite")
	db := database.Connect()

	defer db.Close()
}
