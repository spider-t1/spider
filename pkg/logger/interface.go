package logger

import (
	"context"
	"go.uber.org/zap"
)

// ILogger 统一的日志接口
type ILogger interface {
	// 基础日志方法
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)

	// 带上下文的日志方法
	DebugWithContext(ctx context.Context, msg string, fields ...zap.Field)
	InfoWithContext(ctx context.Context, msg string, fields ...zap.Field)
	WarnWithContext(ctx context.Context, msg string, fields ...zap.Field)
	ErrorWithContext(ctx context.Context, msg string, fields ...zap.Field)
	FatalWithContext(ctx context.Context, msg string, fields ...zap.Field)

	// 业务日志方法
	LogBusiness(ctx context.Context, event BusinessEvent)
	LogAudit(ctx context.Context, event AuditEvent)
	LogAccess(ctx context.Context, event AccessEvent)
	LogSlow(ctx context.Context, event SlowEvent)
}

// BusinessEvent 业务事件
type BusinessEvent struct {
	Type       string                 `json:"type"`        // 业务类型：order, user, payment等
	Action     string                 `json:"action"`      // 操作类型：create, update, delete等
	ResourceID string                 `json:"resource_id"` // 资源ID
	Message    string                 `json:"message"`     // 事件描述
	Data       map[string]interface{} `json:"data"`        // 业务数据
	Result     string                 `json:"result"`      // 操作结果：success, failed
	ErrorCode  string                 `json:"error_code"`  // 错误码
	ErrorMsg   string                 `json:"error_msg"`   // 错误信息
}

// AuditEvent 审计事件
type AuditEvent struct {
	Action     string                 `json:"action"`      // 操作类型
	Resource   string                 `json:"resource"`    // 操作资源
	ResourceID string                 `json:"resource_id"` // 资源ID
	OldValue   map[string]interface{} `json:"old_value"`   // 修改前的值
	NewValue   map[string]interface{} `json:"new_value"`   // 修改后的值
	Result     string                 `json:"result"`      // 操作结果
	Reason     string                 `json:"reason"`      // 操作原因
	IP         string                 `json:"ip"`          // 操作IP
	UserAgent  string                 `json:"user_agent"`  // 用户代理
}

// AccessEvent 访问事件
type AccessEvent struct {
	Method     string `json:"method"`      // HTTP方法
	Cost       int64  `json:"cost"`        // 请求耗时(毫秒)
	Path       string `json:"path"`        // 请求路径
	Query      string `json:"query"`       // 查询参数
	Body       string `json:"body"`        // 请求体
	StatusCode int    `json:"status_code"` // 响应状态码
	ClientIP   string `json:"client_ip"`   // 客户端IP
	UserAgent  string `json:"user_agent"`  // 用户代理
	Referer    string `json:"referer"`     // 来源页面
}

// SlowEvent 慢请求事件
type SlowEvent struct {
	Method     string                 `json:"method"`      // HTTP方法
	Path       string                 `json:"path"`        // 请求路径
	Type       string                 `json:"type"`        // 慢请求类型：api, sql, redis等
	Operation  string                 `json:"operation"`   // 操作名称
	Cost       int64                  `json:"cost"`        // 耗时(毫秒)
	Threshold  int64                  `json:"threshold"`   // 阈值(毫秒)
	Query      string                 `json:"query"`       // 查询语句
	Parameters map[string]interface{} `json:"parameters"`  // 参数
	StackTrace string                 `json:"stack_trace"` // 堆栈信息
	StatusCode int                    `json:"status_code"` // 响应状态码
	ClientIP   string                 `json:"client_ip"`   // 客户端IP
}

// LogContext 日志上下文信息
type LogContext struct {
	RequestID   string `json:"request_id"`   // 请求ID
	UserID      int64  `json:"user_id"`      // 用户ID
	TenantID    int64  `json:"tenant_id"`    // 租户ID
	UserName    string `json:"user_name"`    // 用户名
	TraceID     string `json:"trace_id"`     // 链路追踪ID
	SpanID      string `json:"span_id"`      // Span ID
	ServiceName string `json:"service_name"` // 服务名
	Version     string `json:"version"`      // 版本号
	Environment string `json:"environment"`  // 环境
	Module      string `json:"module"`       // 模块名
	Function    string `json:"function"`     // 函数名
}
