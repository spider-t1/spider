package buyin

import (
    "github.com/gin-gonic/gin"
    buyin_logic "spider/internal/app/logic/buyin"
    "spider/internal/app/response"
    "spider/internal/app/types/types_buyin"
)

// ProductListHandler 商品列表
// @Summary 获取商品列表
// @Description 通过多维度参数与鉴权获取商品列表
// @Tags buyin
// @Produce json
// @Param dateType query string false "日期类型"
// @Param beginDate query string false "开始时间"
// @Param endDate query string false "结束时间"
// @Param isActivity query boolean false "是否活动"
// @Param activityId query string false "活动ID"
// @Param keyWord query string false "关键词"
// @Param indexSelected query string false "指标选择"
// @Param saleType query int false "售卖类型"
// @Param contentType query int false "内容类型"
// @Param cateIds query string false "类目ID列表"
// @Param cateIdsOriginal query int false "原始类目ID"
// @Param productTab query int false "商品标签"
// @Param onlyAbnormal query boolean false "仅异常"
// @Param onlyDropGmv query boolean false "仅GMV下降"
// @Param onlyDropProductShow query boolean false "仅商品曝光下降"
// @Param useCustomizeGmv query boolean false "自定义GMV"
// @Param useCustomizeProductShow query boolean false "自定义商品曝光"
// @Param abnormalThresholdGmv query string false "GMV异常阈值"
// @Param abnormalThresholdProductShow query int false "曝光异常阈值"
// @Param newVersion query boolean false "新版本"
// @Param pageNo query int false "页码"
// @Param pageSize query int false "页大小"
// @Param lid query string false "_lid"
// @Param cookie query string true "Cookie"
// @Param verifyFp query string true "verifyFp"
// @Param fp query string true "fp"
// @Param msToken query string true "msToken"
// @Param ewid query string false "ewid"
// @Param userAgent query string false "User-Agent"
// @Success 200 {object} map[string]interface{}
// @Router /api/buyin/product/list [get]
func ProductListHandler(c *gin.Context) {
    var (
        err error
        req types_buyin.ProductListReq
        src interface{}
    )
    defer func() { response.HandleDefault(c, response.WithSourceData(src))(&err, recover()) }()
    if err = c.ShouldBind(&req); err != nil { return }
    req.Adjust()
    client := buyin_logic.NewJXClient(buyin_logic.ClientConfig{Cookie: req.Cookie, EWID: req.EWID, VerifyFp: req.VerifyFp, Fp: req.Fp, MsToken: req.MsToken, UserAgent: req.UserAgent})
    s1, s2, err := client.GetProductList(buyin_logic.ProductListRequest{DateType: req.DateType, BeginDate: req.BeginDate, EndDate: req.EndDate, IsActivity: req.IsActivity, ActivityID: req.ActivityID, KeyWord: req.KeyWord, IndexSelected: req.IndexSelected, SaleType: req.SaleType, ContentType: req.ContentType, CateIds: req.CateIds, CateIdsOriginal: req.CateIdsOriginal, ProductTab: req.ProductTab, OnlyAbnormal: req.OnlyAbnormal, OnlyDropGmv: req.OnlyDropGmv, OnlyDropProductShow: req.OnlyDropProductShow, UseCustomizeGmv: req.UseCustomizeGmv, UseCustomizeProductShow: req.UseCustomizeProductShow, AbnormalThresholdGmv: req.AbnormalThresholdGmv, AbnormalThresholdProductShow: req.AbnormalThresholdProductShow, NewVersion: req.NewVersion, PageNo: req.PageNo, PageSize: req.PageSize, LID: req.LID})
    if err != nil { return }
    src = map[string]string{"status": s1, "data": s2}
}