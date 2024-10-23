package resolver

import (
	"context"
	"microGo/pkg/common/kafka"
	"time"
)

type Resolver struct {
	kafkaProducer *kafka.Producer
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	return &User{
		ID:        id,
		Email:     "user@example.com",
		Nickname:  "dummy_user",
		Bio:       "This is a dummy user",
		Avatar:    "https://example.com/avatar.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (r *queryResolver) SearchUsers(ctx context.Context, query string, limit *int, offset *int) ([]*User, error) {
	return []*User{
		{
			ID:       "1",
			Email:    "user1@example.com",
			Nickname: "user1",
		},
		{
			ID:       "2",
			Email:    "user2@example.com",
			Nickname: "user2",
		},
	}, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) UpdateProfile(ctx context.Context, input UpdateProfileInput) (*User, error) {
	updatedUser := &User{
		ID:       "1",
		Nickname: input.Nickname,
		Bio:      input.Bio,
	}

	event := UserUpdatedEvent{
		UserID:   "1",
		Nickname: input.Nickname,
		Bio:      input.Bio,
	}

	if err := r.kafkaProducer.Publish("user.updated", event); err != nil {
		return nil, err
	}

	return updatedUser, nil
}
