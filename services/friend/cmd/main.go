package handler

import (
	"github.com/gin-gonic/gin"
	"microGo/pkg/common/kafka"
	"net/http"
)

type FriendHandler struct {
	kafkaProducer *kafka.Producer
}

func NewFriendHandler(producer *kafka.Producer) *FriendHandler {
	return &FriendHandler{
		kafkaProducer: producer,
	}
}

func (h *FriendHandler) SendFriendRequest(c *gin.Context) {
	var req struct {
		ToUserID string `json:"to_user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := FriendRequestEvent{
		FromUserID: c.GetString("user_id"),
		ToUserID:   req.ToUserID,
		Status:     "pending",
	}

	if err := h.kafkaProducer.Publish("friend.request.created", event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
