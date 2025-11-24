package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// SquareFilterHandler 达人广场筛选条件
// @Summary 获取达人广场筛选条件
// @Description 通过 type/reqScene 与鉴权参数获取广场筛选条件
// @Tags buyin
// @Produce json
// @Param type query string false "类型"
// @Param reqScene query string false "场景"
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Success 200 {object} buyin_logic.SquareFilterFile
// @Router /api/buyin/square/filter [get]
func SquareFilterHandler(c *gin.Context) {
    var (
        err error
        req types_buyin.SquareFilterReq
        res = &buyin_logic.SquareFilterFile{}
    )
    defer func() { response.HandleDefault(c, response.WithData(res))(&err, recover()) }()
    if err = c.ShouldBind(&req); err != nil { return }
    req.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: req.Cookie, EWID: req.EWID, VerifyFp: req.VerifyFp, Fp: req.Fp, MsToken: req.MsToken, UserAgent: req.UserAgent})
    var httpStatus string
    res, httpStatus, err = client.SquareFilter(&buyin_logic.SquareFilterParams{Type: req.Type, ReqScene: req.ReqScene})
    _ = httpStatus
    if err != nil { return }
}