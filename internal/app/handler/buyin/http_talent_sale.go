package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// TalentSaleAnalyzeHandler 达人销售分析
// @Summary 获取达人销售分析
// @Description 通过 uid 与鉴权参数获取达人销售分析
// @Tags buyin
// @Produce json
// @Param uid query string true "达人UID（v2_开头）"
// @Param range query string false "时间范围，如30d，默认30d"
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Success 200 {object} buyin_logic.TalentSaleAnalyzeResp
// @Router /api/buyin/author/sales [get]
func TalentSaleAnalyzeHandler(c *gin.Context) {
    var (
        err error
        req types_buyin.TalentSaleAnalyzeReq
        res = &buyin_logic.TalentSaleAnalyzeResp{}
    )
    defer func() { response.HandleDefault(c, response.WithData(res))(&err, recover()) }()
    if err = c.ShouldBind(&req); err != nil { return }
    req.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: req.Cookie, EWID: req.EWID, VerifyFp: req.VerifyFp, Fp: req.Fp, MsToken: req.MsToken, UserAgent: req.UserAgent})
    var status string
    res, status, err = client.TalentSaleAnaly(buyin_logic.TalentSaleAnalyzeParams{UID: req.UID, Range: req.Range})
    _ = status
    if err != nil { return }
}