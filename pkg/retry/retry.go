package retry

import (
	"github.com/cenkalti/backoff/v4"
	"time"
)

type RetryableFunc func() error

type RetryConfig struct {
	MaxRetries      int
	InitialInterval time.Duration
	MaxInterval     time.Duration
	Multiplier      float64
	MaxElapsedTime  time.Duration
}

func DefaultConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:      3,
		InitialInterval: 100 * time.Millisecond,
		MaxInterval:     10 * time.Second,
		Multiplier:      2.0,
		MaxElapsedTime:  1 * time.Minute,
	}
}

func WithRetry(fn RetryableFunc, cfg RetryConfig) error {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = cfg.InitialInterval
	b.MaxInterval = cfg.MaxInterval
	b.Multiplier = cfg.Multiplier
	b.MaxElapsedTime = cfg.MaxElapsedTime

	return backoff.Retry(backoff.Operation(fn), b)
}
