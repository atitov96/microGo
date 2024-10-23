package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"microGo/pkg/health"
	"microGo/pkg/metrics"
	"microGo/pkg/telemetry"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Printf("%e", err)
			return
		}
	}(logger)

	cleanup, err := telemetry.InitTracer(telemetry.TracerConfig{
		ServiceName:    "your-service",
		ServiceVersion: "1.0.0",
		Environment:    os.Getenv("ENV"),
		JaegerURL:      os.Getenv("JAEGER_URL"),
	})
	if err != nil {
		logger.Fatal("failed to init tracer", zap.Error(err))
	}
	defer cleanup()

	db, _ := sql.Open("", "")

	//kafkaConfig := kafka.Config{
	//	Brokers:     []string{os.Getenv("KAFKA_BROKERS")},
	//	RetryConfig: retry.DefaultConfig(),
	//	DLQTopic:    "dlq",
	//	RetryTopic:  "retry",
	//}
	//kafkaClient, err := kgo.NewClient(kafkaConfig)
	seeds := []string{"localhost:9092"}
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup("my-group-identifier"),
		kgo.ConsumeTopics("foo"),
	)
	if err != nil {
		logger.Fatal("failed to create kafka client", zap.Error(err))
	}

	newMetrics := metrics.NewMetrics("your-service")

	healthier := health.NewHealth(newMetrics)
	healthier.AddChecker(&health.DatabaseChecker{DB: db})
	healthier.AddChecker(&health.KafkaChecker{Client: kafkaClient})

	router := chi.NewRouter()

	router.Use(
		telemetry.TracingMiddleware,
		metrics.Middleware(newMetrics),
		middleware.Recoverer,
	)

	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/health", healthier.Handler())

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("server shutdown failed", zap.Error(err))
		}
	}()

	logger.Info("server starting", zap.String("addr", server.Addr))
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Fatal("server failed", zap.Error(err))
	}
}
