package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"spider/internal/middleware/metadata"
	"spider/pkg/logger"
	"strings"
	"time"
)

// AccessLogMiddleware 访问日志中间件
// 记录所有请求的基本信息，包括请求ID
func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 读取请求体内容
		var dataStr string
		if c.Request.Body != nil {
			data, err := io.ReadAll(c.Request.Body)
			if err != nil {
				logger.ErrorWithContext(c.Request.Context(), "获取请求体内容失败", zap.Error(err))
				dataStr = ""
			} else {
				dataStr = string(data)
			}
			// 重新设置请求体，以便后续处理器可以再次读取
			c.Request.Body = io.NopCloser(strings.NewReader(dataStr))
		}

		// 使用新的结构化访问日志

		metadata.SetMetadataForRequestId(c)
		logger.LogAccess(c.Request.Context(), logger.AccessEvent{
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Query:     c.Request.URL.RawQuery,
			Body:      strings.ReplaceAll(dataStr, "\r\n", ""),
			ClientIP:  c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Referer:   c.Request.Referer(),
		})

		// 处理请求
		c.Next()
		duration := time.Since(start)

		// 记录请求结束
		logger.LogAccess(c.Request.Context(), logger.AccessEvent{
			Cost:       duration.Milliseconds(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			Query:      c.Request.URL.RawQuery,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			UserAgent:  c.Request.UserAgent(),
			Referer:    c.Request.Referer(),
		})
	}
}
