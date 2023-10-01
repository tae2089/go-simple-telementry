package main

import (
	"context"
	"net/http"
	"telemetry/common/config"
	"telemetry/common/telemetry"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/bob-logging/logger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {

	telemetryConfig := config.GetTelemetryConfig()
	shutdown, err := telemetry.SetupOTelSDK(context.Background(), telemetryConfig.ServiceName, telemetryConfig.ServiceVersion)
	if err != nil {
		logger.Error(err)
	}
	defer shutdown(context.Background())
	r := gin.Default()
	r.Use(otelgin.Middleware(telemetryConfig.ServiceName))
	r.GET("/", hello)
	r.GET("/health", healthCheck)
	r.Run()
}

func healthCheck(c *gin.Context) {
	logger.Info("health check")
	c.String(http.StatusOK, "health")
}

func hello(c *gin.Context) {
	logger.Info("hello")
	c.String(http.StatusOK, "hello")
}
