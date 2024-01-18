package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// default route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! two",
		})
	})

	router.Run(":8082")
}
