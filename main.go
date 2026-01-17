package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Golang Gin API",
			"status":  "running",
		})
	})

	// Define a health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Define a sample API endpoint
	r.GET("/api/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	// Start the server on port 8080
	r.Run(":8080")
}
