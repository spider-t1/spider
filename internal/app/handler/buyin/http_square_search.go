package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// DarenSquareSearchHandler 达人广场搜索
// @Summary 搜索达人广场
// @Description Query 传鉴权参数，Body 传搜索条件
// @Tags buyin
// @Produce json
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Param body body buyin_logic.DarenSquareRequest true "搜索条件"
// @Success 200 {object} buyin_logic.FeedAuthorList
// @Router /api/buyin/square/search [post]
func DarenSquareSearchHandler(c *gin.Context) {
    var (
        err error
        auth types_buyin.BuyinAuthReq
        body buyin_logic.DarenSquareRequest
        res = &buyin_logic.FeedAuthorList{}
    )
    defer func() { response.HandleDefault(c, response.WithData(res))(&err, recover()) }()
    if err = c.ShouldBindQuery(&auth); err != nil { return }
    if err = c.ShouldBindJSON(&body); err != nil { return }
    auth.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: auth.Cookie, EWID: auth.EWID, VerifyFp: auth.VerifyFp, Fp: auth.Fp, MsToken: auth.MsToken, UserAgent: auth.UserAgent})
    res, err = client.SearchFeedAuthor(body)
    if err != nil { return }
}