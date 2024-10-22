package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendMessageRequest struct {
	ToUser  string `json:"to_user" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	r.GET("/ready", readinessCheck)

	messages := r.Group("/messages")
	{
		messages.POST("", sendMessage)

		messages.GET("/:chatId", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "get chat messages"})
		})
	}

	r.Run(":8080")
}

func sendMessage(c *gin.Context) {
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement actual message sending logic
	c.JSON(http.StatusOK, gin.H{
		"message": "Message sent",
		"to_user": req.ToUser,
		"content": req.Content,
	})
}

func readinessCheck(context *gin.Context) {
	// TODO: Check if the database is ready
	context.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func healthCheck(с *gin.Context) {
	с.JSON(http.StatusOK, gin.H{"status": "alive"})
}
