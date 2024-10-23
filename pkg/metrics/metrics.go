package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"time"
)

type Metrics struct {
	RequestCounter   *prometheus.CounterVec
	RequestDuration  *prometheus.HistogramVec
	MessageProcessed *prometheus.CounterVec
	QueueSize        *prometheus.GaugeVec
	ErrorCounter     *prometheus.CounterVec
}

func NewMetrics(serviceName string) *Metrics {
	return &Metrics{
		RequestCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name:        "http_requests_total",
				Help:        "Total number of HTTP requests",
				ConstLabels: prometheus.Labels{"service": serviceName},
			},
			[]string{"method", "endpoint", "status"},
		),

		RequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:        "http_request_duration_seconds",
				Help:        "HTTP request duration in seconds",
				ConstLabels: prometheus.Labels{"service": serviceName},
				Buckets:     []float64{0.1, 0.3, 0.5, 0.7, 1, 3, 5, 7, 10},
			},
			[]string{"method", "endpoint"},
		),

		MessageProcessed: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name:        "kafka_messages_processed_total",
				Help:        "Total number of Kafka messages processed",
				ConstLabels: prometheus.Labels{"service": serviceName},
			},
			[]string{"topic", "status"},
		),

		QueueSize: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name:        "kafka_queue_size",
				Help:        "Current size of Kafka message queue",
				ConstLabels: prometheus.Labels{"service": serviceName},
			},
			[]string{"topic"},
		),

		ErrorCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name:        "error_total",
				Help:        "Total number of errors",
				ConstLabels: prometheus.Labels{"service": serviceName},
			},
			[]string{"type"},
		),
	}
}

func Middleware(metrics *Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			//wrapped := wrapResponseWriter(w)
			wrapped := w

			next.ServeHTTP(wrapped, r)

			duration := time.Since(start).Seconds()
			//metrics.RequestCounter.WithLabelValues(r.Method, r.URL.Path, fmt.Sprint(wrapped.status)).Inc()
			metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
		})
	}
}
