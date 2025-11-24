package router

import (
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "spider/docs" // main 文件中导入 docs 包
    "spider/internal/config"
    "spider/internal/middleware"
    "spider/internal/router/buyin"
    "spider/internal/router/douyin"
    "time"
)

func InitRouter(r *gin.Engine) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    UseMiddleware(r)

    // 注册抖音相关路由
    douyin.InitDouyinRoute(r)
    buyin.InitBuyinRoute(r)
}

func UseMiddleware(r *gin.Engine) {

	// 使用请求ID中间件（必须在其他中间件之前）
	r.Use(middleware.RequestID())
	// 使用访问日志中间件
	r.Use(middleware.AccessLogMiddleware())
	// 使用慢日志中间件，阈值设置为 3 秒
	r.Use(middleware.SlowLogMiddleware(time.Duration(config.Cfg.Logging.Categories.Slow.Threshold) * time.Second))
	r.Use(middleware.CorsMiddleware())
}
