package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}

func main() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	r.GET("/ready", readinessCheck)

	users := r.Group("/users")
	{
		users.GET("/:id", getProfile)
		users.PUT("/:id", updateProfile)
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

func getProfile(c *gin.Context) {
	userID := c.Param("id")

	// TODO: Implement actual profile retrieval
	c.JSON(http.StatusOK, gin.H{
		"id":       userID,
		"nickname": "example_user",
		"bio":      "Example bio",
		"avatar":   "example_avatar_url",
	})
}

func updateProfile(c *gin.Context) {
	userID := c.Param("id")
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement actual profile update
	c.JSON(http.StatusOK, gin.H{
		"id":       userID,
		"nickname": req.Nickname,
		"bio":      req.Bio,
		"avatar":   req.Avatar,
	})
}
