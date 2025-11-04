package logger

import (
	"context"
	"go.uber.org/zap"
)

// LogBusiness 记录业务日志
func LogBusiness(ctx context.Context, event BusinessEvent) {
	if Logger == nil {
		return
	}

	// 构建基础字段
	fields := BuildContextFields(ctx)

	// 添加业务事件字段
	businessFields := BuildBusinessFields(event)
	fields = append(fields, businessFields...)

	// 记录日志
	Logger.Info(event.Message, fields...)
}

// LogAudit 记录审计日志
func LogAudit(ctx context.Context, event AuditEvent) {
	if Logger == nil {
		return
	}

	// 构建基础字段
	fields := BuildContextFields(ctx)

	// 添加审计事件字段
	auditFields := BuildAuditFields(event)
	fields = append(fields, auditFields...)

	// 记录日志
	message := "audit: " + event.Action + " " + event.Resource
	Logger.Info(message, fields...)
}

// LogAccess 记录访问日志
func LogAccess(ctx context.Context, event AccessEvent) {
	if Logger == nil {
		return
	}

	fields := make([]zap.Field, 0)
	fields = append(fields, zap.String("api", event.Method))
	fields = append(fields, zap.Int64("cost", event.Cost))
	// 构建基础字段
	_fields := BuildContextFields(ctx)
	fields = append(fields, _fields...)
	// 添加访问事件字段
	accessFields := BuildAccessFields(event)
	fields = append(fields, accessFields...)

	// 记录日志
	//message := "access: " + event.Method + " " + event.Path
	InfoWithContext(ctx, "", fields...)
}

// LogSlow 记录慢请求日志
func LogSlow(ctx context.Context, event SlowEvent) {
	if Logger == nil {
		return
	}
	fields := make([]zap.Field, 0)
	fields = append(fields, zap.String("api", "slow"))
	fields = append(fields, zap.Int64("cost", event.Cost))
	// 构建基础字段
	_fields := BuildContextFields(ctx)
	fields = append(fields, _fields...)
	// 添加慢请求事件字段
	slowFields := BuildSlowFields(event)
	fields = append(fields, slowFields...)

	// 记录日志
	//message := "slow: " + event.Type + " " + event.Operation
	Logger.Warn("", fields...)
}

// 便利方法

// LogUserAction 记录用户操作
func LogUserAction(ctx context.Context, action, resource, resourceID string, data map[string]interface{}) {
	LogBusiness(ctx, BusinessEvent{
		Type:       "user_action",
		Action:     action,
		ResourceID: resourceID,
		Message:    "用户操作: " + action + " " + resource,
		Data:       data,
		Result:     "success",
	})
}

// LogUserActionError 记录用户操作错误
func LogUserActionError(ctx context.Context, action, resource, resourceID string, err error, errorCode string) {
	LogBusiness(ctx, BusinessEvent{
		Type:       "user_action",
		Action:     action,
		ResourceID: resourceID,
		Message:    "用户操作失败: " + action + " " + resource,
		Result:     "failed",
		ErrorCode:  errorCode,
		ErrorMsg:   err.Error(),
	})
}

// LogDataChange 记录数据变更审计
func LogDataChange(ctx context.Context, resource, resourceID string, oldValue, newValue map[string]interface{}, reason string) {
	LogAudit(ctx, AuditEvent{
		Action:     "update",
		Resource:   resource,
		ResourceID: resourceID,
		OldValue:   oldValue,
		NewValue:   newValue,
		Result:     "success",
		Reason:     reason,
	})
}

// LogDataCreate 记录数据创建审计
func LogDataCreate(ctx context.Context, resource, resourceID string, value map[string]interface{}) {
	LogAudit(ctx, AuditEvent{
		Action:     "create",
		Resource:   resource,
		ResourceID: resourceID,
		NewValue:   value,
		Result:     "success",
	})
}

// LogDataDelete 记录数据删除审计
func LogDataDelete(ctx context.Context, resource, resourceID string, value map[string]interface{}, reason string) {
	LogAudit(ctx, AuditEvent{
		Action:     "delete",
		Resource:   resource,
		ResourceID: resourceID,
		OldValue:   value,
		Result:     "success",
		Reason:     reason,
	})
}

// LogSlowAPI 记录慢API
func LogSlowAPI(ctx context.Context, path string, cost, threshold int64, parameters map[string]interface{}) {
	LogSlow(ctx, SlowEvent{
		Type:       "api",
		Operation:  path,
		Cost:       cost,
		Threshold:  threshold,
		Parameters: parameters,
	})
}

// LogSlowSQL 记录慢SQL
func LogSlowSQL(ctx context.Context, query string, cost, threshold int64, parameters map[string]interface{}) {
	LogSlow(ctx, SlowEvent{
		Type:       "sql",
		Operation:  "database_query",
		Cost:       cost,
		Threshold:  threshold,
		Query:      query,
		Parameters: parameters,
	})
}

// LogSlowRedis 记录慢Redis操作
func LogSlowRedis(ctx context.Context, operation string, cost, threshold int64, parameters map[string]interface{}) {
	LogSlow(ctx, SlowEvent{
		Type:       "redis",
		Operation:  operation,
		Cost:       cost,
		Threshold:  threshold,
		Parameters: parameters,
	})
}
