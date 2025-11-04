package logger

import (
	"context"
	"runtime"
	"spider/internal/config"
	"spider/internal/middleware/metadata"
	"strings"

	"go.uber.org/zap"
)

// ExtractLogContext 从上下文中提取日志相关信息
func ExtractLogContext(ctx context.Context) LogContext {
	logCtx := LogContext{
		RequestID:   metadata.GetRequestId(ctx),
		UserID:      metadata.GetUserId(ctx),
		TenantID:    metadata.GetTenantId(ctx),
		UserName:    metadata.GetUserName(ctx),
		ServiceName: "spider",              // 可以从配置中读取
		Environment: config.Cfg.System.Env, // 可以从配置中读取
	}

	// 尝试从上下文中获取其他信息
	if traceID := getFromContext(ctx, "trace_id"); traceID != "" {
		logCtx.TraceID = traceID
	}
	if spanID := getFromContext(ctx, "span_id"); spanID != "" {
		logCtx.SpanID = spanID
	}
	if version := getFromContext(ctx, "version"); version != "" {
		logCtx.Version = version
	}

	return logCtx
}

// BuildContextFields 构建上下文相关的日志字段
func BuildContextFields(ctx context.Context) []zap.Field {
	logCtx := ExtractLogContext(ctx)
	fields := make([]zap.Field, 0, 10)

	if logCtx.RequestID != "" {
		fields = append(fields, zap.String("request_id", logCtx.RequestID))
	}
	if logCtx.UserID > 0 {
		fields = append(fields, zap.Int64("user_id", logCtx.UserID))
	}
	if logCtx.TenantID > 0 {
		fields = append(fields, zap.Int64("tenant_id", logCtx.TenantID))
	}
	if logCtx.UserName != "" {
		fields = append(fields, zap.String("user_name", logCtx.UserName))
	}
	if logCtx.TraceID != "" {
		fields = append(fields, zap.String("trace_id", logCtx.TraceID))
	}
	if logCtx.SpanID != "" {
		fields = append(fields, zap.String("span_id", logCtx.SpanID))
	}
	if logCtx.ServiceName != "" {
		fields = append(fields, zap.String("service", logCtx.ServiceName))
	}
	if logCtx.Environment != "" {
		fields = append(fields, zap.String("env", logCtx.Environment))
	}
	if logCtx.Version != "" {
		fields = append(fields, zap.String("version", logCtx.Version))
	}

	return fields
}

// BuildBusinessFields 构建业务事件相关的日志字段
func BuildBusinessFields(event BusinessEvent) []zap.Field {
	fields := []zap.Field{
		zap.String("event_type", "business"),
		zap.String("business_type", event.Type),
		zap.String("action", event.Action),
		zap.String("result", event.Result),
	}

	if event.ResourceID != "" {
		fields = append(fields, zap.String("resource_id", event.ResourceID))
	}
	if event.ErrorCode != "" {
		fields = append(fields, zap.String("error_code", event.ErrorCode))
	}
	if event.ErrorMsg != "" {
		fields = append(fields, zap.String("error_msg", event.ErrorMsg))
	}
	if len(event.Data) > 0 {
		fields = append(fields, zap.Any("data", event.Data))
	}

	return fields
}

// BuildAuditFields 构建审计事件相关的日志字段
func BuildAuditFields(event AuditEvent) []zap.Field {
	fields := []zap.Field{
		zap.String("event_type", "audit"),
		zap.String("action", event.Action),
		zap.String("resource", event.Resource),
		zap.String("result", event.Result),
	}

	if event.ResourceID != "" {
		fields = append(fields, zap.String("resource_id", event.ResourceID))
	}
	if event.Reason != "" {
		fields = append(fields, zap.String("reason", event.Reason))
	}
	if event.IP != "" {
		fields = append(fields, zap.String("ip", event.IP))
	}
	if event.UserAgent != "" {
		fields = append(fields, zap.String("user_agent", event.UserAgent))
	}
	if len(event.OldValue) > 0 {
		fields = append(fields, zap.Any("old_value", event.OldValue))
	}
	if len(event.NewValue) > 0 {
		fields = append(fields, zap.Any("new_value", event.NewValue))
	}

	return fields
}

// BuildAccessFields 构建访问事件相关的日志字段
func BuildAccessFields(event AccessEvent) []zap.Field {
	fields := []zap.Field{
		//zap.Int64("cost", event.Cost),
		zap.String("event_type", "access"),
		zap.String("method", event.Method),
		zap.String("path", event.Path),
		zap.Int("status_code", event.StatusCode),
		zap.String("client_ip", event.ClientIP),
	}

	if event.Query != "" {
		fields = append(fields, zap.String("query", event.Query))
	}
	if event.Body != "" {
		fields = append(fields, zap.String("body", event.Body))
	}
	if event.UserAgent != "" {
		fields = append(fields, zap.String("user_agent", event.UserAgent))
	}
	if event.Referer != "" {
		fields = append(fields, zap.String("referer", event.Referer))
	}

	return fields
}

// BuildSlowFields 构建慢请求事件相关的日志字段
func BuildSlowFields(event SlowEvent) []zap.Field {
	fields := []zap.Field{
		zap.String("event_type", "slow"),
		zap.String("slow_type", event.Type),
		zap.String("operation", event.Operation),
		//zap.Int64("duration", event.Cost),
		zap.Int64("threshold", event.Threshold),
	}

	if event.Query != "" {
		fields = append(fields, zap.String("query", event.Query))
	}
	if len(event.Parameters) > 0 {
		fields = append(fields, zap.Any("parameters", event.Parameters))
	}
	if event.StackTrace != "" {
		fields = append(fields, zap.String("stack_trace", event.StackTrace))
	}

	return fields
}

// GetCallerInfo 获取调用者信息
func GetCallerInfo(skip int) (string, string) {
	pc, file, _, ok := runtime.Caller(skip)
	if !ok {
		return "", ""
	}

	// 获取函数名
	funcName := runtime.FuncForPC(pc).Name()
	if idx := strings.LastIndex(funcName, "/"); idx != -1 {
		funcName = funcName[idx+1:]
	}

	// 获取文件名
	if idx := strings.LastIndex(file, "/"); idx != -1 {
		file = file[idx+1:]
	}

	return funcName, file
}

// getFromContext 从上下文中获取字符串值
func getFromContext(ctx context.Context, key string) string {
	if value := ctx.Value(key); value != nil {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}
