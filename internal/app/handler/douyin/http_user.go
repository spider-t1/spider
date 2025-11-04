package douyin

import (
	"github.com/gin-gonic/gin"
	douyin2 "spider/internal/app/logic/douyin"
	"spider/internal/app/response"
	"spider/internal/app/types/types_douyin"
)

// UserInfoHandler 获取抖音用户信息
// @Summary 获取抖音用户信息
// @Description 通过 secId 获取用户资料
// @Tags douyin
// @Produce json
// @Param secId query string true "用户sec_id"
// @Success 200 {object} types_douyin.DouyinUserInfoResp "返回JSON结构"
// @Router /api/douyin/user/info [get]
func UserInfoHandler(c *gin.Context) {

	var (
		err error
		ctx = c.Request.Context()
		req types_douyin.UserInfoReq
		res = &types_douyin.DouyinUserInfoResp{}
	)

	defer func() {
		response.HandleDefault(c, response.WithData(res))(&err, recover())
	}()

	if err = c.ShouldBind(&req); err != nil {
		return
	}
	req.Adjust()

	client := douyin2.NewDouyinClient("")
	res, err = client.DouyinUserInfo(ctx, req.SecId)
	if err != nil {
		return
	}
}

// UserSearchHandler 搜索抖音用户
// @Summary 搜索抖音用户
// @Description 通过 keyword 搜索用户
// @Tags douyin
// @Produce json
// @Param keyword query string true "搜索关键词"
// @Success 200 {object} types_douyin.DouyinSearchResp "返回JSON结构"
// @Router /api/douyin/user/search [get]
func UserSearchHandler(c *gin.Context) {

	var (
		err error
		ctx = c.Request.Context()
		req types_douyin.UserSearchReq
		res = &types_douyin.DouyinSearchResp{}
	)

	defer func() {
		response.HandleDefault(c, response.WithData(res))(&err, recover())
	}()

	if err = c.ShouldBind(&req); err != nil {
		return
	}

	req.Adjust()

	client := douyin2.NewDouyinClient("")
	res, err = client.DouyinUserSearch(ctx, req.Keyword)
	if err != nil {
		return
	}
}

// UserVideoHandler 获取抖音用户作品列表
// @Summary 获取抖音用户作品列表
// @Description 通过 secId 获取用户作品列表
// @Tags douyin
// @Produce json
// @Param secId query string true "用户secId"
// @Success 200 {object} types_douyin.DyOneVideoInfo "返回原始JSON"
// @Router /api/douyin/user/video [get]
func UserVideoHandler(c *gin.Context) {

	var (
		err error
		ctx = c.Request.Context()
		req types_douyin.UserVideoReq
		res = &types_douyin.DyOneVideoInfo{}
	)

	defer func() {
		response.HandleDefault(c, response.WithData(res))(&err, recover())
	}()

	if err = c.ShouldBind(&req); err != nil {
		return
	}
	req.Adjust()

	client := douyin2.NewDouyinClient("")
	res, err = client.DouyinUserVideo(ctx, req.SecId)
	if err != nil {
		return
	}

}
