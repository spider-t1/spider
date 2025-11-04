package middleware

import (
	"github.com/gin-gonic/gin"
	"spider/pkg/logger"
	"time"
)

func SlowLogMiddleware(threshold time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		duration := time.Since(start)

		// 如果超过阈值就记录慢日志
		if duration > threshold {
			logger.LogSlow(c.Request.Context(), logger.SlowEvent{
				Type:       "api",
				Cost:       duration.Milliseconds(),
				Method:     c.Request.Method,
				Path:       c.Request.URL.Path,
				Threshold:  threshold.Milliseconds(),
				StatusCode: c.Writer.Status(),
				ClientIP:   c.ClientIP(),
			})
		}
	}
}
