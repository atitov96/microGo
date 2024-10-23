package health

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"microGo/pkg/metrics"
	"net/http"
	"sync"
	"time"
)

type Status string

const (
	StatusUp   Status = "UP"
	StatusDown Status = "DOWN"
)

type Checker interface {
	Name() string
	Check(context.Context) error
}

type Check struct {
	Status    Status    `json:"status"`
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type Health struct {
	mu       sync.RWMutex
	checkers []Checker
	metrics  *metrics.Metrics
}

func NewHealth(metrics *metrics.Metrics) *Health {
	return &Health{
		checkers: make([]Checker, 0),
		metrics:  metrics,
	}
}

func (h *Health) AddChecker(checker Checker) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.checkers = append(h.checkers, checker)
}

func (h *Health) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		status := StatusUp
		checks := make(map[string]Check)

		for _, checker := range h.checkers {
			check := Check{
				Status:    StatusUp,
				Timestamp: time.Now(),
			}

			if err := checker.Check(ctx); err != nil {
				check.Status = StatusDown
				check.Message = err.Error()
				status = StatusDown
				h.metrics.ErrorCounter.WithLabelValues("health_check").Inc()
			}

			checks[checker.Name()] = check
		}

		response := struct {
			Status string           `json:"status"`
			Checks map[string]Check `json:"checks"`
		}{
			Status: string(status),
			Checks: checks,
		}

		w.Header().Set("Content-Type", "application/json")
		if status == StatusDown {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Printf("%e", err)
			return
		}
	}
}

type DatabaseChecker struct {
	DB *sql.DB
}

func (c *DatabaseChecker) Name() string {
	return "database"
}

func (c *DatabaseChecker) Check(ctx context.Context) error {
	return c.DB.PingContext(ctx)
}

type KafkaChecker struct {
	Client *kgo.Client
}

func (c *KafkaChecker) Name() string {
	return "kafka"
}

func (c *KafkaChecker) Check(ctx context.Context) error {
	return c.Client.Ping(ctx)
}
