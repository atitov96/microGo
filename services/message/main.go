package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	r.GET("/ready", readinessCheck)

	messages := r.Group("/messages")
	{
		messages.POST("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "send message"})
		})

		messages.GET("/:chatId", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "get chat messages"})
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
