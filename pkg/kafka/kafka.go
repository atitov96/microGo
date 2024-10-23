package kafka

import (
	"context"
	"encoding/json"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

type Consumer struct {
	client *kgo.Client
}

type Message struct {
	Topic   string          `json:"topic"`
	Key     string          `json:"key"`
	Payload json.RawMessage `json:"payload"`
}

func NewProducer(brokers []string) (*Producer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		return nil, err
	}

	return &Producer{client: client}, nil
}

func NewConsumer(brokers []string, groupID string) (*Consumer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(groupID),
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{client: client}, nil
}

func (p *Producer) Publish(topic string, key string, value interface{}) error {
	payload, err := json.Marshal(value)
	if err != nil {
		return err
	}

	record := &kgo.Record{
		Topic: topic,
		Key:   []byte(key),
		Value: payload,
	}

	return p.client.ProduceSync(context.Background(), record).FirstErr()
}

func (c *Consumer) Subscribe(topics []string, handler func(Message) error) error {
	//c.client.SetTopics(topics...)

	for {
		fetches := c.client.PollFetches(context.Background())
		if fetches.IsClientClosed() {
			return nil
		}

		//for _, fetch := range fetches {
		//	for _, record := range fetch.Records {
		//		msg := Message{
		//			Topic:   record.Topic,
		//			Key:     string(record.Key),
		//			Payload: record.Value,
		//		}
		//
		//		if err := handler(msg); err != nil {
		//			return err
		//		}
		//	}
		//}
	}
}
