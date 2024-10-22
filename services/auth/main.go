package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	r.GET("/ready", readinessCheck)

	auth := r.Group("/auth")
	{
		auth.POST("/register", register)
		auth.POST("/login", login)
		auth.POST("/oauth", oauthLogin)
	}

	r.Run(":8080")
}

func oauthLogin(c *gin.Context) {
	// TODO: Implement OAuth login
	c.JSON(http.StatusOK, gin.H{"message": "oauth endpoint"})
}

func login(c *gin.Context) {
	// TODO: Implement login
	c.JSON(http.StatusOK, gin.H{"message": "login endpoint"})
}

func register(c *gin.Context) {
	// TODO: Implement register
	c.JSON(http.StatusOK, gin.H{"message": "register endpoint"})
}

func readinessCheck(context *gin.Context) {
	// TODO: Check if the database is ready
	context.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func healthCheck(с *gin.Context) {
	с.JSON(http.StatusOK, gin.H{"status": "alive"})
}
