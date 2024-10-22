package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	r.GET("/ready", readinessCheck)

	users := r.Group("/users")
	{
		users.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "get user profile"})
		})
		users.PUT("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "update user profile"})
		})
		users.GET("/search", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "search users"})
		})
	}

	r.Run(":8080")
}

func readinessCheck(context *gin.Context) {
	// TODO: Check if the database is ready
	context.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func healthCheck(с *gin.Context) {
	с.JSON(http.StatusOK, gin.H{"status": "alive"})
}
