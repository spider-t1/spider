package buyin

import (
    "github.com/gin-gonic/gin"
    h "spider/internal/app/handler/buyin"
)

func InitBuyinRoute(r *gin.Engine) {
    g := r.Group("/api/buyin")
    g.GET("/shop/user", h.ShopUserInfoHandler)
    g.GET("/dashboard/operating", h.FrontPageOperatingDataHandler)
    g.GET("/author/fans", h.FansAnalyzeHandler)
    g.GET("/author/sales", h.TalentSaleAnalyzeHandler)
    g.GET("/author/profile", h.AllianceProfileHandler)
    g.GET("/square/filter", h.SquareFilterHandler)
    g.POST("/square/search", h.DarenSquareSearchHandler)
    g.GET("/live/overview", h.LiveListOverviewHandler)
    g.GET("/video/list", h.VideoListHandler)
    g.GET("/product/list", h.ProductListHandler)
}