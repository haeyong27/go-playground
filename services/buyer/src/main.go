package main

import (
	"riad/package/aws"
)

// func main() {
// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Initialize the database

// 	database := datasource.NewDatabase("m")
// 	database.Connect()

// 	r := adapter.NewRouter()
// 	s := adapter.NewServer(r)
// 	s.Routing()
// 	s.Router.Run()

// 	// Start the server
// 	router.Run(":8080")
// }

func main() {
	for {
		aws.ReceiveMessage()
	}
}
