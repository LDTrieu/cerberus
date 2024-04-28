package gzlog

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

type ContextFn func(ctx context.Context) []zapcore.Field

type Logger struct {
	LogLevel      gormlogger.LogLevel
	SlowThreshold time.Duration
}

func New(level gormlogger.LogLevel) Logger {
	return Logger{
		LogLevel:      level,
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l Logger) SetAsDefault() {
	gormlogger.Default = l
}

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return Logger{
		SlowThreshold: l.SlowThreshold,
		LogLevel:      level,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}

	ctxzap.Extract(ctx).Sugar().Infof(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}

	ctxzap.Extract(ctx).Sugar().Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}

	ctxzap.Extract(ctx).Sugar().Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	logger := ctxzap.Extract(ctx)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error:
		sql, rows := fc()
		logger.Error("[TRACE] ", zap.Error(err), zap.Duration("duration", elapsed), zap.Int64("rows", rows), zap.String("sql query", sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		logger.Warn("[TRACE] ", zap.Duration("duration", elapsed), zap.Int64("rows", rows), zap.String("sql query", sql))
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		logger.Info("[TRACE] ", zap.Duration("duration", elapsed), zap.Int64("rows", rows), zap.String("sql query", sql))
	}
}
