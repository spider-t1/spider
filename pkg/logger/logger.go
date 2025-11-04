package logger

import (
	"context"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"spider/internal/config"
	"spider/internal/middleware/metadata"
	"strconv"
	"strings"
	"time"
)

const (
	colorDebug = "\033[36m" // Cyan
	colorInfo  = "\033[32m" // Green
	colorWarn  = "\033[33m" // Yellow
	colorError = "\033[31m" // Red
	colorReset = "\033[0m"  // Reset
)

// 自定义时间编码器 - 输出格式：年-月-日 时:分:秒
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

var Logger *zap.Logger

// InitLogger 初始化日志系统
func InitLogger() error {
	return InitLoggerWithConfig(config.Cfg.Logging)
}

// InitLoggerWithConfig 使用配置初始化日志系统
func InitLoggerWithConfig(cfg config.Logging) error {
	// 设置默认配置
	if cfg.Level == "" {
		cfg.Level = "info"
	}
	if cfg.Format == "" {
		cfg.Format = "json"
	}

	// 解析日志级别
	level, err := parseLogLevel(cfg.Level)
	if err != nil {
		return fmt.Errorf("invalid log level %s: %w", cfg.Level, err)
	}

	var cores []zapcore.Core

	// 处理输出配置
	for _, output := range cfg.Output {
		core, err := createCore(output, cfg.Format, level)
		if err != nil {
			return fmt.Errorf("failed to create core for output %s: %w", output.Type, err)
		}
		cores = append(cores, core)
	}

	// 处理分类日志
	if cfg.Categories.Access.Enabled {
		core, err := createFileCore(cfg.Categories.Access.Path, cfg.Format, zapcore.InfoLevel)
		if err != nil {
			return fmt.Errorf("failed to create access log core: %w", err)
		}
		cores = append(cores, core)
	}

	if cfg.Categories.Error.Enabled {
		core, err := createFileCore(cfg.Categories.Error.Path, cfg.Format, zapcore.ErrorLevel)
		if err != nil {
			return fmt.Errorf("failed to create error log core: %w", err)
		}
		cores = append(cores, core)
	}

	if cfg.Categories.Slow.Enabled {
		warnOnly := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.WarnLevel
		})
		core, err := createFileCore(cfg.Categories.Slow.Path, cfg.Format, warnOnly)
		if err != nil {
			return fmt.Errorf("failed to create slow log core: %w", err)
		}
		cores = append(cores, core)
	}

	// 如果没有配置任何输出，使用默认控制台输出
	if len(cores) == 0 {
		consoleCore, err := createConsoleCore(cfg.Format, level, true)
		if err != nil {
			return fmt.Errorf("failed to create default console core: %w", err)
		}
		cores = append(cores, consoleCore)
	}

	// 创建logger
	core := zapcore.NewTee(cores...)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(Logger)

	return nil
}

func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// 带颜色的级别输出
	encoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		var color string
		switch level {
		case zapcore.DebugLevel:
			color = colorDebug
		case zapcore.InfoLevel:
			color = colorInfo
		case zapcore.WarnLevel:
			color = colorWarn
		case zapcore.ErrorLevel:
			color = colorError
		default:
			color = colorReset
		}
		enc.AppendString(color + level.CapitalString() + colorReset)
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

//func getEncoder() zapcore.Encoder {
//	cfg := zap.NewProductionEncoderConfig()
//	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
//	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
//	cfg.EncodeCaller = zapcore.ShortCallerEncoder
//	return zapcore.NewJSONEncoder(cfg)
//}

// InfoWithContext 带请求ID的便利日志函数
func InfoWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	requestID := metadata.GetRequestId(ctx)
	if requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
		fields = append(fields, zap.String("msg", msg))
	}
	Logger.Info("", fields...)
}

func InfoFWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	requestID := metadata.GetRequestId(ctx)
	if requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
		fields = append(fields, zap.String("msg", msg))
	}
	Logger.Info("", fields...)
}

func WarnWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	requestID := metadata.GetRequestId(ctx)
	if requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}
	Logger.Warn(msg, fields...)
}

func ErrorWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	requestID := metadata.GetRequestId(ctx)
	if requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}
	Logger.Error(msg, fields...)
}

func DebugWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	requestID := metadata.GetRequestId(ctx)
	if requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}
	Logger.Debug(msg, fields...)
}

// parseLogLevel 解析日志级别
func parseLogLevel(level string) (zapcore.Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn", "warning":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	case "fatal":
		return zapcore.FatalLevel, nil
	default:
		return zapcore.InfoLevel, fmt.Errorf("unknown log level: %s", level)
	}
}

// parseDuration 解析时间字符串
func parseDuration(s string) (time.Duration, error) {
	if s == "" {
		return 0, fmt.Errorf("empty duration string")
	}

	// 处理天数
	if strings.HasSuffix(s, "d") {
		days, err := strconv.Atoi(strings.TrimSuffix(s, "d"))
		if err != nil {
			return 0, err
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}

	// 使用标准时间解析
	return time.ParseDuration(s)
}

// createCore 创建日志核心
func createCore(output config.LogOutput, format string, level zapcore.LevelEnabler) (zapcore.Core, error) {
	switch output.Type {
	//case "file":
	//return createFileCore(output.Path, format, level)
	case "console":
		return createConsoleCore(format, level, output.Colored)
	default:
		return nil, fmt.Errorf("unsupported output type: %s", output.Type)
	}
}

// createFileCore 创建文件日志核心
func createFileCore(path, format string, level zapcore.LevelEnabler) (zapcore.Core, error) {
	if path == "" {
		return nil, fmt.Errorf("file output path is required")
	}

	// 确保目录存在
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory %s: %w", path, err)
	}

	// 解析轮转配置
	maxAge := 30 * 24 * time.Hour // 默认30天
	if config.Cfg.Logging.MaxAge != "" {
		var err error
		maxAge, err = parseDuration(config.Cfg.Logging.MaxAge)
		if err != nil {
			return nil, fmt.Errorf("invalid max_age %s: %w", config.Cfg.Logging.MaxAge, err)
		}
	}

	rotationTime := 24 * time.Hour // 默认24小时
	if config.Cfg.Logging.RotationTime != "" {
		var err error
		rotationTime, err = parseDuration(config.Cfg.Logging.RotationTime)
		if err != nil {
			return nil, fmt.Errorf("invalid rotation_time %s: %w", config.Cfg.Logging.RotationTime, err)
		}
	}

	// 创建轮转日志
	logPath := filepath.Join(path, "app_%Y-%m-%d.log")
	writer, err := rotatelogs.New(
		logPath,
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create rotatelogs writer: %w", err)
	}

	// 文件输出始终使用JSON格式，便于日志分析
	encoder := getFileEncoder()
	return zapcore.NewCore(encoder, zapcore.AddSync(writer), level), nil
}

// createConsoleCore 创建控制台日志核心
func createConsoleCore(format string, level zapcore.LevelEnabler, colored bool) (zapcore.Core, error) {
	var encoder zapcore.Encoder
	// 控制台输出始终使用带颜色的文本格式，便于开发时查看
	if colored {
		encoder = getConsoleEncoder()
	} else {
		encoder = getTextEncoder()
	}

	return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level), nil
}

// getEncoder 根据格式获取编码器
func getEncoder(format string) zapcore.Encoder {
	switch format {
	case "json":
		return getFileEncoder()
	case "text":
		return getTextEncoder()
	default:
		return getFileEncoder()
	}
}

// getTextEncoder 获取文本编码器
func getTextEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// ZapLogger 实现ILogger接口的结构体
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger 创建新的ZapLogger实例
func NewZapLogger(logger *zap.Logger) ILogger {
	return &ZapLogger{logger: logger}
}

// 实现ILogger接口的基础日志方法
func (l *ZapLogger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *ZapLogger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *ZapLogger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

// 实现ILogger接口的带上下文日志方法
func (l *ZapLogger) DebugWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := BuildContextFields(ctx)
	allFields := append(contextFields, fields...)
	l.logger.Debug(msg, allFields...)
}

func (l *ZapLogger) InfoWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := BuildContextFields(ctx)
	allFields := append(contextFields, fields...)
	l.logger.Info(msg, allFields...)
}

func (l *ZapLogger) WarnWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := BuildContextFields(ctx)
	allFields := append(contextFields, fields...)
	l.logger.Warn(msg, allFields...)
}

func (l *ZapLogger) ErrorWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := BuildContextFields(ctx)
	allFields := append(contextFields, fields...)
	l.logger.Error(msg, allFields...)
}

func (l *ZapLogger) FatalWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	contextFields := BuildContextFields(ctx)
	allFields := append(contextFields, fields...)
	l.logger.Fatal(msg, allFields...)
}

// 实现ILogger接口的业务日志方法
func (l *ZapLogger) LogBusiness(ctx context.Context, event BusinessEvent) {
	LogBusiness(ctx, event)
}

func (l *ZapLogger) LogAudit(ctx context.Context, event AuditEvent) {
	LogAudit(ctx, event)
}

func (l *ZapLogger) LogAccess(ctx context.Context, event AccessEvent) {
	LogAccess(ctx, event)
}

func (l *ZapLogger) LogSlow(ctx context.Context, event SlowEvent) {
	LogSlow(ctx, event)
}

// GetLogger 获取ILogger实例
func GetLogger() ILogger {
	if Logger == nil {
		return nil
	}
	return NewZapLogger(Logger)
}
