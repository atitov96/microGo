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
		friends.POST("/request", sendFriendRequest)

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

func sendFriendRequest(c *gin.Context) {
	var req struct {
		ToUserID string `json:"to_user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement actual friend request logic
	c.JSON(http.StatusOK, gin.H{
		"message":    "Friend request sent",
		"to_user_id": req.ToUserID,
	})
}

func readinessCheck(context *gin.Context) {
	// TODO: Check if the database is ready
	context.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func healthCheck(с *gin.Context) {
	с.JSON(http.StatusOK, gin.H{"status": "alive"})
}
