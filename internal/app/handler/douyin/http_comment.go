package douyin

import (
	"github.com/gin-gonic/gin"
	douyin2 "spider/internal/app/logic/douyin"
	"spider/internal/app/response"
	"spider/internal/app/types/types_douyin"
)

// CommentListHandler 获取抖音评论列表
// @Summary 获取抖音评论列表
// @Description 通过 aweme_id 获取评论列表
// @Tags douyin
// @Produce json
// @Param awemeId query string true "视频ID，默认7561795410549853481"
// @Param cursor query string false "分页游标，默认0"
// @Param count query string false "每页数量，默认10"
// @Success 200 {object} map[string]interface{} "返回原始JSON"
// @Router /api/douyin/comment/list [get]
func CommentListHandler(c *gin.Context) {

	var (
		err error
		ctx = c.Request.Context()
		req types_douyin.CommentListReq
		res = &types_douyin.CommentData{}
	)

	defer func() {
		response.HandleDefault(c, response.WithData(res))(&err, recover())
	}()

	if err = c.ShouldBind(&req); err != nil {
		return
	}
	req.Adjust()

	client := douyin2.NewDouyinClient("")
	res, err = client.DouyinComment(ctx, &req)
	if err != nil {
		return
	}
}

// CommentReplyListHandler 获取抖音评论回复列表
// @Summary 获取抖音评论回复列表
// @Description 通过 item_id 与 comment_id 获取评论回复列表
// @Tags douyin
// @Produce json
// @Param item_id query string true "作品ID"
// @Param comment_id query string true "评论ID"
// @Param cursor query string false "分页游标，默认0"
// @Param count query string false "每页数量，默认10"
// @Success 200 {object} map[string]interface{} "返回原始JSON"
// @Router /api/douyin/comment/reply [get]
func CommentReplyListHandler(c *gin.Context) {

	var (
		err error
		ctx = c.Request.Context()
		req types_douyin.CommentReplyListReq
		res = &types_douyin.CommentData{}
	)

	defer func() {
		response.HandleDefault(c, response.WithData(res))(&err, recover())
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}
	req.Adjust()

	client := douyin2.NewDouyinClient("")
	res, err = client.DouyinCommentReply(ctx, &req)
	if err != nil {
		return
	}
}
