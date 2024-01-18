package main

import (
	"buyer/src/adapter"
	"buyer/src/datasource"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Initialize the database
	database := datasource.NewDatabase("m")
	database.Connect()

	r := adapter.NewRouter()
	s := adapter.NewServer(r)
	s.Run()

	// Start the server
	router.Run(":8080")
}
