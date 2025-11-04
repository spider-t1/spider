package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"spider/internal/middleware/metadata"
)

type ZapGormLogger struct {
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

func NewGormLogger(threshold time.Duration) logger.Interface {
	return &ZapGormLogger{
		SlowThreshold: threshold,
		LogLevel:      logger.Info, // 你可以控制 gorm 默认日志输出级别
	}
}

func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	requestID := metadata.GetRequestId(ctx)
	if requestID != "" {
		Logger.Info(msg, zap.String("request_id", requestID), zap.Any("data", data))
	} else {
		Logger.Sugar().Infof(msg, data...)
	}
}

func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	requestID := metadata.GetRequestId(ctx)
	if requestID != "" {
		Logger.Warn(msg, zap.String("request_id", requestID), zap.Any("data", data))
	} else {
		Logger.Sugar().Warnf(msg, data...)
	}
}

func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	requestID := metadata.GetRequestId(ctx)
	if requestID != "" {
		Logger.Error(msg, zap.String("request_id", requestID), zap.Any("data", data))
	} else {
		Logger.Sugar().Errorf(msg, data...)
	}
}

func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zap.Field{}

	fields = append(fields, zap.String("sql", "exec"))
	fields = append(fields, zap.Int64("cost", elapsed.Milliseconds()))

	if err != nil && l.LogLevel >= logger.Error {
		fields = append(fields, zap.String("err", err.Error()))
	}

	requestID := metadata.GetRequestId(ctx)
	// 添加请求ID到字段中
	if requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}

	fields = append(fields,
		//zap.Duration("cost", elapsed),
		zap.Int64("rows", rows),
		zap.String("sql", sql),
		zap.String("file", utils.FileWithLineNum()),
	)

	// 错误 SQL -> error.log
	if err != nil && l.LogLevel >= logger.Error {
		Logger.Error("", fields...)
	}

	// 慢 SQL -> slow.log
	if elapsed > l.SlowThreshold {
		Logger.Warn("慢 SQL", fields...)
	}

	// 所有 SQL -> app.log
	if l.LogLevel >= logger.Info {
		Logger.Info("", fields...)
	}

}
