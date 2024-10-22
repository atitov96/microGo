package kafka

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaClient struct {
	_ *kgo.Client
}

func NewKafkaClient() {

}

func (k *KafkaClient) PublishMessage(topic string, message interface{}) error {
	return nil
}

func (k *KafkaClient) ConsumeMessage(topic string, handler func([]byte) error) error {
	return nil
}
