package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// ShopUserInfoHandler 获取精选联盟账号信息
// @Summary 获取精选联盟账号信息
// @Description 通过鉴权参数读取当前精选联盟账号信息
// @Tags buyin
// @Produce json
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Success 200 {object} buyin_logic.ShopUserInfoResp
// @Router /api/buyin/shop/user [get]
func ShopUserInfoHandler(c *gin.Context) {
    var (
        err error
        req types_buyin.ShopUserInfoReq
        res = &buyin_logic.ShopUserInfoResp{}
    )
    defer func() { response.HandleDefault(c, response.WithData(res))(&err, recover()) }()
    if err = c.ShouldBind(&req); err != nil { return }
    req.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: req.Cookie, EWID: req.EWID, VerifyFp: req.VerifyFp, Fp: req.Fp, MsToken: req.MsToken, UserAgent: req.UserAgent})
    res, err = client.GetShopUserInfo()
    if err != nil { return }
}