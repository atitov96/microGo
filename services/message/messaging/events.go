package messaging

import (
	"github.com/gin-gonic/gin"
	"time"
)

type MessageEvent struct {
	Type      string    `json:"type"`
	MessageID string    `json:"message_id"`
	FromUser  string    `json:"from_user"`
	ToUser    string    `json:"to_user"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func sendMessage(c *gin.Context) {

}
