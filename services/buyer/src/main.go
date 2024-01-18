package main

import (
	"buyer/src/adapter"
	"buyer/src/datasource"
)

func main() {
	// Create a new Gin router
	// router := gin.Default()

	// Initialize the database
	database := datasource.NewDatabase("m")
	database.Connect()

	r := adapter.NewRouter()
	s := adapter.NewServer(r)
	s.Routing()
	s.Router.Run()

	// Start the server
	// router.Run(":8080")
}
