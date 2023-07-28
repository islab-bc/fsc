package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = "stacktrace"
	config.EncoderConfig = encoderConfig
	config.OutputPaths = []string{"stdout", "/tmp/user.log"}
	log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

// Info ...
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug ...
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error ...
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
