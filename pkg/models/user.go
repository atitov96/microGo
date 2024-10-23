package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // never return password
	Nickname string `json:"nickname"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}

type FriendRequest struct {
	ID         string `json:"id"`
	FromUserID string `json:"from_user_id"`
	ToUserID   string `json:"to_user_id"`
	Status     string `json:"status"` // pending, accepted, rejected
	CreatedAt  string `json:"created_at"`
}

type Message struct {
	ID        string `json:"id"`
	FromUser  string `json:"from_user"`
	ToUser    string `json:"to_user"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
