package handler

import (
	"github.com/gin-gonic/gin"
	"microGo/pkg/kafka"
	"time"
)

type FriendHandler struct {
	producer   *kafka.Producer
	userClient UserServiceClient
}

func (h *FriendHandler) SendFriendRequest(c *gin.Context) {
	var req struct {
		ToUserID string `json:"to_user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.userClient.GetUser(c.Request.Context(), req.ToUserID); err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	requestID := generateID()
	fromUserID := c.GetString("user_id")

	// Публикуем событие
	event := &events.FriendRequestEvent{
		ID:         requestID,
		FromUserID: fromUserID,
		ToUserID:   req.ToUserID,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	if err := h.producer.Publish(events.TopicFriendRequest, requestID, event); err != nil {
		c.JSON(500, gin.H{"error": "failed to publish event"})
		return
	}

	c.JSON(200, event)
}

func (h *FriendHandler) StartUserUpdatesConsumer() error {
	consumer, err := kafka.NewConsumer([]string{"localhost:9092"}, "friend-service")
	if err != nil {
		return err
	}

	return consumer.Subscribe([]string{events.TopicUserUpdated}, func(msg kafka.Message) error {
		var event events.UserUpdatedEvent
		if err := json.Unmarshal(msg.Payload, &event); err != nil {
			return err
		}

		return h.updateUserCache(event)
	})
}
