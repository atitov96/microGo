package kafka

import "time"

// User Service Events

type UserCreatedEvent struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdatedEvent struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Bio       string    `json:"bio"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Friend Service Events

type FriendRequestEvent struct {
	ID         string    `json:"id"`
	FromUserID string    `json:"from_user_id"`
	ToUserID   string    `json:"to_user_id"`
	Status     string    `json:"status"` // pending, accepted, rejected
	CreatedAt  time.Time `json:"created_at"`
}

type FriendshipStatusChangedEvent struct {
	FromUserID string    `json:"from_user_id"`
	ToUserID   string    `json:"to_user_id"`
	Status     string    `json:"status"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Message Service Events

type MessageSentEvent struct {
	ID        string    `json:"id"`
	FromUser  string    `json:"from_user"`
	ToUser    string    `json:"to_user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	TopicUserCreated       = "user.created"
	TopicUserUpdated       = "user.updated"
	TopicFriendRequest     = "friend.request"
	TopicFriendshipChanged = "friend.status.changed"
	TopicMessageSent       = "message.sent"
)
