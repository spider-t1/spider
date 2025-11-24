package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// FansAnalyzeHandler 达人粉丝分析
// @Summary 获取达人粉丝分析
// @Description 通过 uid 与鉴权参数获取达人粉丝分析
// @Tags buyin
// @Produce json
// @Param uid query string true "达人UID（v2_开头）"
// @Param fansClub query string false "粉丝团参数，默认0"
// @Param worksType query string false "内容类型，默认1"
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Success 200 {object} buyin_logic.FansAnalyzeResp
// @Router /api/buyin/author/fans [get]
func FansAnalyzeHandler(c *gin.Context) {
    var (
        err error
        req types_buyin.FansAnalyzeReq
        res = &buyin_logic.FansAnalyzeResp{}
    )
    defer func() { response.HandleDefault(c, response.WithData(res))(&err, recover()) }()
    if err = c.ShouldBind(&req); err != nil { return }
    req.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: req.Cookie, EWID: req.EWID, VerifyFp: req.VerifyFp, Fp: req.Fp, MsToken: req.MsToken, UserAgent: req.UserAgent})
    res, err = client.FansAnalyze(&buyin_logic.FansAnalyzeParams{UID: req.UID, FansClub: req.FansClub, WorksType: req.WorksType})
    if err != nil { return }
}