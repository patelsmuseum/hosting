package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Starting server...")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.Run(":8080")
}
