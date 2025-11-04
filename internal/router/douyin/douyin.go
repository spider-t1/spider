package douyin

import (
	"github.com/gin-gonic/gin"
	h "spider/internal/app/handler/douyin"
)

func InitDouyinRoute(r *gin.Engine) {
    g := r.Group("/api/douyin")
    g.GET("/comment/list", h.CommentListHandler)
    g.GET("/comment/reply", h.CommentReplyListHandler)
    g.GET("/aweme/detail", h.AwemeDetailHandler)
    g.GET("/user/info", h.UserInfoHandler)
    g.GET("/user/search", h.UserSearchHandler)
    g.GET("/user/video", h.UserVideoHandler)
}
