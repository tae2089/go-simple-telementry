package main

import (
	"context"
	"net/http"
	"time"

	"github.com/tae2089/go-simple-telemetry/common/config"
	"github.com/tae2089/go-simple-telemetry/common/telemetry"
	"github.com/tae2089/go-simple-telemetry/log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

var (
	latencyHistogramExamplar prometheus.Histogram = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_request_latency_examplar",
			Help:    "Latency of HTTP requests examplar",
			Buckets: prometheus.DefBuckets,
		},
		// []string{"path", "method", "traceId"},
	)
	latencyHistogram *prometheus.HistogramVec = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_latency",
			Help:    "Latency of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

func main() {
	telemetryConfig := config.GetTelemetryConfig()
	shutdown, err := telemetry.SetupOTelSDK(context.Background(), telemetryConfig.ServiceName, telemetryConfig.ServiceVersion)
	if err != nil {
		log.Error(err.Error())
	}
	defer shutdown(context.Background())
	r := gin.New()

	r.GET("/", hello)
	r.GET("/health", healthCheck)
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	)))
	r.Run()
}

func healthCheck(c *gin.Context) {
	log.Info("health check")
	c.String(http.StatusOK, "health")
}

func hello(c *gin.Context) {
	start := time.Now()
	_, span := otel.GetTracerProvider().Tracer("hello api").Start(c.Request.Context(), "hello")
	span.SetAttributes(attribute.String("path", c.Request.URL.Path))
	span.SetAttributes(attribute.String("method", c.Request.Method))
	log.Info("hello", zap.String("traceId", span.SpanContext().TraceID().String()))
	defer span.End()
	elapsed := time.Since(start)
	measureLatency(c.Request.URL.Path, c.Request.Method, span.SpanContext().TraceID().String(), elapsed)
	c.String(http.StatusOK, "hello")
}

func measureLatency(path string, method string, traceId string, duration time.Duration) {
	if examplar, ok := latencyHistogramExamplar.(prometheus.ExemplarObserver); ok {
		examplar.ObserveWithExemplar(duration.Seconds(), prometheus.Labels{"path": path, "method": method, "traceId": traceId})
	}
	latencyHistogram.With(prometheus.Labels{"path": path, "method": method}).Observe(duration.Seconds())
}
