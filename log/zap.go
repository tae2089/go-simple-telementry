package log

import (
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}
	logger = zap.Must(config.Build(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.TimeKey = "timestamp"
		// encoderConfig.MessageKey = zapcore.OmitKey
		encoderConfig.LevelKey = zapcore.OmitKey
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder := zapcore.NewJSONEncoder(encoderConfig)
		stdout := zapcore.AddSync(os.Stdout)
		core := zapcore.NewTee(
			zapcore.NewCore(encoder, stdout, zapcore.DebugLevel),
		)
		return core
	})))

}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Errorf(err error, fields ...zap.Field) {
	logger.Error(err.Error(), fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func GetLoggingFieldsByRequest(r *http.Request, headers []string, statusCode int, duration time.Duration) []zap.Field {
	var query string = ""
	if r.URL.RawQuery != "" {
		query = "?" + r.URL.RawQuery
	}
	loggingDefault := []zap.Field{
		zap.Int("Status Code", statusCode),
		zap.String("Remote-Host", r.RemoteAddr),
		zap.String("Method", r.Method),
		zap.String("Path", r.URL.Path+query),
		zap.Duration("Response Time", duration),
	}
	loggingHeaders := createLoggingHeaders(r, headers)
	loggingFields := append(loggingDefault, loggingHeaders...)
	return loggingFields
}

func createLoggingHeaders(r *http.Request, headers []string) []zap.Field {
	var loggingHeaders []zap.Field
	loggingHeaders = append(loggingHeaders, zap.String("TraceID", r.Header.Get("Traceparent")))
	for _, header := range headers {
		loggingHeaders = append(loggingHeaders, zap.String(header, r.Header.Get(header)))
	}
	return loggingHeaders
}
