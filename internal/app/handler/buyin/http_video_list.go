package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// VideoListHandler 视频列表
// @Summary 获取视频列表
// @Description 通过多维度参数与鉴权获取视频列表
// @Tags buyin
// @Produce json
// @Param pageNo query string false "页码"
// @Param pageSize query string false "页大小"
// @Param beginDate query string false "开始时间"
// @Param endDate query string false "结束时间"
// @Param dateType query string false "日期类型"
// @Param activityId query string false "活动ID"
// @Param accountType query string false "账号类型"
// @Param authorId query string false "作者ID"
// @Param rangeType query string false "范围类型"
// @Param cartType query string false "挂车类型"
// @Param adType query string false "广告类型"
// @Param searchInfo query string false "搜索信息"
// @Param indexSelected query string false "指标选择"
// @Param sortField query string false "排序字段"
// @Param lid query string false "_lid"
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Success 200 {object} map[string]interface{}
// @Router /api/buyin/video/list [get]
func VideoListHandler(c *gin.Context) {
    var (
        err error
        req types_buyin.VideoListReq
        raw string
    )
    defer func() { response.HandleDefault(c, response.WithSourceData(raw))(&err, recover()) }()
    if err = c.ShouldBind(&req); err != nil { return }
    req.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: req.Cookie, EWID: req.EWID, VerifyFp: req.VerifyFp, Fp: req.Fp, MsToken: req.MsToken, UserAgent: req.UserAgent})
    raw, _, err = client.VideoList(buyin_logic.VideoListParams{PageNo: req.PageNo, PageSize: req.PageSize, BeginDate: req.BeginDate, EndDate: req.EndDate, DateType: req.DateType, ActivityID: req.ActivityID, AccountType: req.AccountType, AuthorID: req.AuthorID, RangeType: req.RangeType, CartType: req.CartType, AdType: req.AdType, SearchInfo: req.SearchInfo, IndexSelected: req.IndexSelected, SortField: req.SortField, LID: req.LID})
    if err != nil { return }
}