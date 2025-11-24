package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// FrontPageOperatingDataHandler 首页运营数据
// @Summary 获取首页运营数据
// @Description 通过 timeRange 与鉴权参数获取首页运营数据
// @Tags buyin
// @Produce json
// @Param timeRange query string true "时间范围"
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Success 200 {object} map[string]interface{}
// @Router /api/buyin/dashboard/operating [get]
func FrontPageOperatingDataHandler(c *gin.Context) {
    var (
        err error
        req types_buyin.FrontPageOperatingDataReq
        raw string
    )
    defer func() { response.HandleDefault(c, response.WithSourceData(raw))(&err, recover()) }()
    if err = c.ShouldBind(&req); err != nil { return }
    req.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: req.Cookie, EWID: req.EWID, VerifyFp: req.VerifyFp, Fp: req.Fp, MsToken: req.MsToken, UserAgent: req.UserAgent})
    raw, _, err = client.GetFrontPageOperatingData(req.TimeRange)
    if err != nil { return }
}