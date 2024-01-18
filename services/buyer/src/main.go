package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Initialize the database
	// database := datasource.NewDatabase("m")
	// database.Connect()

	// Add your routes here
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// router.GET("/users", func(c *gin.Context) {
	// 	users := query.GetUser()
	// 	c.JSON(200, gin.H{
	// 		"users": users,
	// 	})
	// })

	// router.GET("/user", func(c *gin.Context) {
	// 	user := query.CreateUser()
	// 	c.JSON(200, gin.H{
	// 		"user": user,
	// 	})
	// })

	// Start the server
	router.Run(":8080")
}
