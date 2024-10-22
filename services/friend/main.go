package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	r.GET("/ready", readinessCheck)

	friends := r.Group("/friends")
	{
		friends.POST("/request", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "send friend request"})
		})

		friends.PUT("/request/:id/accept", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "accept friend request"})
		})

		friends.PUT("/request/:id/reject", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "reject friend request"})
		})

		friends.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "get friends list"})
		})

		friends.DELETE("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "remove friend"})
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
