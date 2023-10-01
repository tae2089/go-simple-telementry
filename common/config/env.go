package config

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/tae2089/bob-logging/logger"
)

type TelemetryConfig struct {
	ServiceName    string `env:"SERVICE_NAME,required"`
	ServiceVersion string `env:"SERVICE_VERSION,required"`
}

var (
	telemetryConfig = &TelemetryConfig{}
)

func init() {
	environment := os.Getenv("ENVIRONMENT")

	if environment == "local" || environment == "development" {
		projectDir := getProjectDir() + "/.env"
		logger.Info(projectDir)
		err := godotenv.Load(projectDir)
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	if err := env.Parse(telemetryConfig); err != nil {
		panic(err)
	}
}

func getProjectDir() string {
	projectDir := ""
	_, filename, _, _ := runtime.Caller(0)
	projectDir = path.Join(path.Dir(filename), "../..")
	return projectDir
}

func GetTelemetryConfig() *TelemetryConfig {
	return telemetryConfig
}
