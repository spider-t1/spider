package douyin

import (
	"github.com/gin-gonic/gin"
	douyin2 "spider/internal/app/logic/douyin"
	"spider/internal/app/response"
	"spider/internal/app/types/types_douyin"
)

// AwemeDetailHandler 获取抖音作品详情
// @Summary 获取抖音作品详情
// @Description 通过 aweme_id 获取作品详情
// @Tags douyin
// @Produce json
// @Param aweme_id query string true "作品ID"
// @Success 200 {object} map[string]interface{} "返回原始JSON"
// @Router /api/douyin/aweme/detail [get]
func AwemeDetailHandler(c *gin.Context) {

	var (
		err error
		ctx = c.Request.Context()
		req types_douyin.AwemeDetailReq
		res = &types_douyin.DouyinUserInfoResp{}
	)

	defer func() {
		response.HandleDefault(c, response.WithData(res))(&err, recover())
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}
	req.Adjust()

	client := douyin2.NewDouyinClient("")
	body, err := client.DouyinAwemeDetail(ctx, req.AwemeId)
	if err != nil {
		response.HandleDefault(c)(&err, nil)
		return
	}
	response.Success(c, gin.H{"raw": body})
}
