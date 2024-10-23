package service

import (
	"context"
	"log"
	"microGo/pkg/kafka"
	"time"
)

type UserService struct {
	producer *kafka.Producer
}

func (s *UserService) UpdateProfile(ctx context.Context, input UpdateProfileInput) (*User, error) {
	user, err := s.repository.UpdateProfile(ctx, input)
	if err != nil {
		return nil, err
	}

	event := &events.UserUpdatedEvent{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Bio:       user.Bio,
		UpdatedAt: time.Now(),
	}

	if err := s.producer.Publish(events.TopicUserUpdated, user.ID, event); err != nil {
		log.Printf("failed to publish user update event: %v", err)
	}

	return user, nil
}

func (s *UserService) StartFriendshipConsumer() error {
	consumer, err := kafka.NewConsumer([]string{"localhost:9092"}, "user-service")
	if err != nil {
		return err
	}

	return consumer.Subscribe([]string{events.TopicFriendshipChanged}, func(msg kafka.Message) error {
		var event events.FriendshipStatusChangedEvent
		if err := json.Unmarshal(msg.Payload, &event); err != nil {
			return err
		}

		return s.updateFriendshipCounters(event)
	})
}
