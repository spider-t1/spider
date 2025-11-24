package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// LiveListOverviewHandler 直播列表概览
// @Summary 获取直播列表概览
// @Description 通过多维度参数与鉴权获取直播列表概览
// @Tags buyin
// @Produce json
// @Param aType query string false "类型"
// @Param dateType query string false "日期类型"
// @Param beginDate query string false "开始时间"
// @Param endDate query string false "结束时间"
// @Param sortField query string false "排序字段"
// @Param isAsc query string false "升序"
// @Param pageNo query string false "页码"
// @Param pageSize query string false "页大小"
// @Param version query string false "版本"
// @Param lid query string false "_lid"
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Success 200 {object} map[string]interface{}
// @Router /api/buyin/live/overview [get]
func LiveListOverviewHandler(c *gin.Context) {
    var (
        err error
        req types_buyin.LiveListOverviewReq
        raw string
    )
    defer func() { response.HandleDefault(c, response.WithSourceData(raw))(&err, recover()) }()
    if err = c.ShouldBind(&req); err != nil { return }
    req.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: req.Cookie, EWID: req.EWID, VerifyFp: req.VerifyFp, Fp: req.Fp, MsToken: req.MsToken, UserAgent: req.UserAgent})
    raw, _, err = client.LiveListOverview(buyin_logic.LiveListOverviewParams{AType: req.AType, DateType: req.DateType, BeginDate: req.BeginDate, EndDate: req.EndDate, SortField: req.SortField, IsAsc: req.IsAsc, PageNo: req.PageNo, PageSize: req.PageSize, Version: req.Version, LID: req.LID})
    if err != nil { return }
}