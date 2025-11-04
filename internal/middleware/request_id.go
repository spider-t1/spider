package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"spider/internal/app/consts"
	"spider/internal/config"
)

// RequestID 请求ID中间件
// 为每个请求生成唯一的请求ID，并将其存储在上下文中
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试从请求头中获取请求ID
		requestID := c.GetHeader(consts.RequestIDKey)

		// 如果请求头中没有请求ID，则生成一个新的
		if requestID == "" {
			if config.Cfg.System.Env == "dev" {
				requestID = "dev-" + uuid.New().String()
			} else {
				// 正式环境必须要求req_id
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid req_id"})
			}
		}

		// 将请求ID存储在上下文中
		c.Set(consts.RequestIDKey, requestID)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), consts.RequestIDKey, requestID))

		// 将请求ID添加到响应头中
		c.Header(consts.RequestIDKey, requestID)

		// 继续处理请求
		c.Next()
	}
}

// GetRequestID 从Gin上下文中获取请求ID
func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(consts.RequestIDKey); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return ""
}
