package logger

import (
	"os"

	"github.com/ldtrieu/cerberus/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func Init(cfg *config.Logger, opts ...zap.Option) (*zap.Logger, error) {
	var (
		zapConfig zap.Config
		encode    zapcore.LevelEncoder
	)
	

	switch os.Getenv("ENVIRONMENT") {
	case "local":
		encode = zapcore.CapitalColorLevelEncoder
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.Encoding = "console"
	case "staging", "prod":
		encode = zapcore.CapitalLevelEncoder
		zapConfig = zap.NewProductionConfig()
		zapConfig.Encoding = "json"
	default:
		encode = zapcore.CapitalLevelEncoder
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.Encoding = "json"
	}

	switch cfg.Level {
	case "debug":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "warn":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	zapConfig.EncoderConfig.MessageKey = "message"
	zapConfig.EncoderConfig.TimeKey = "ts"
	zapConfig.EncoderConfig.LevelKey = "level"
	zapConfig.EncoderConfig.NameKey = "log"
	if !cfg.DisableCaller {
		zapConfig.EncoderConfig.CallerKey = "caller"
	}
	if !cfg.DisableStacktrace {
		zapConfig.EncoderConfig.StacktraceKey = "stacktrace"
	}
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncoderConfig.EncodeLevel = encode
	zapConfig.Sampling = nil

	return zapConfig.Build(opts...)
}
