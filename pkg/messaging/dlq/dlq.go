package dlq

import (
	"encoding/json"
	"microGo/pkg/kafka"
	"time"
)

type FailedMessage struct {
	OriginalTopic string          `json:"original_topic"`
	Payload       json.RawMessage `json:"payload"`
	Error         string          `json:"error"`
	RetryCount    int             `json:"retry_count"`
	LastRetry     time.Time       `json:"last_retry"`
	OriginalTime  time.Time       `json:"original_time"`
}

type DLQHandler struct {
	producer   kafka.Producer
	dlqTopic   string
	retryTopic string
}

func NewDLQHandler(producer kafka.Producer, dlqTopic, retryTopic string) *DLQHandler {
	return &DLQHandler{
		producer:   producer,
		dlqTopic:   dlqTopic,
		retryTopic: retryTopic,
	}
}

func (h *DLQHandler) HandleFailedMessage(msg FailedMessage) error {
	if msg.RetryCount < 3 {
		msg.RetryCount++
		msg.LastRetry = time.Now()
		return h.producer.Publish(h.retryTopic, msg.OriginalTopic, msg)
	}

	return h.producer.Publish(h.dlqTopic, msg.OriginalTopic, msg)
}
