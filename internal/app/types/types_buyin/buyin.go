package types_buyin

import (
    "spider/internal/app/types/types_common"
)

type BuyinAuthReq struct {
    types_common.BaseParam
    Cookie    string `form:"cookie" binding:"required"`
    EWID      string `form:"ewid"`
    VerifyFp  string `form:"verifyFp" binding:"required"`
    Fp        string `form:"fp" binding:"required"`
    MsToken   string `form:"msToken" binding:"required"`
    UserAgent string `form:"userAgent"`
}

func (r *BuyinAuthReq) Adjust() {
    if r.UserAgent == "" {
        r.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0"
    }
}

type ShopUserInfoReq struct {
    BuyinAuthReq
}

func (r *ShopUserInfoReq) Adjust() {
    r.BuyinAuthReq.Adjust()
}

type FrontPageOperatingDataReq struct {
    BuyinAuthReq
    TimeRange string `form:"timeRange" binding:"required"`
}

func (r *FrontPageOperatingDataReq) Adjust() {
    r.BuyinAuthReq.Adjust()
}

type FansAnalyzeReq struct {
    BuyinAuthReq
    UID       string `form:"uid" binding:"required"`
    FansClub  string `form:"fansClub"`
    WorksType string `form:"worksType"`
}

func (r *FansAnalyzeReq) Adjust() {
    if r.FansClub == "" {
        r.FansClub = "0"
    }
    if r.WorksType == "" {
        r.WorksType = "1"
    }
    r.BuyinAuthReq.Adjust()
}

type AllianceProfileReq struct {
    BuyinAuthReq
    UID       string `form:"uid" binding:"required"`
    WorksType string `form:"worksType"`
}

func (r *AllianceProfileReq) Adjust() {
    if r.WorksType == "" {
        r.WorksType = "1"
    }
    r.BuyinAuthReq.Adjust()
}

type SquareFilterReq struct {
    BuyinAuthReq
    Type     string `form:"type"`
    ReqScene string `form:"reqScene"`
}

func (r *SquareFilterReq) Adjust() {
    r.BuyinAuthReq.Adjust()
}

type LiveListOverviewReq struct {
    BuyinAuthReq
    AType     string `form:"aType"`
    DateType  string `form:"dateType"`
    BeginDate string `form:"beginDate"`
    EndDate   string `form:"endDate"`
    SortField string `form:"sortField"`
    IsAsc     string `form:"isAsc"`
    PageNo    string `form:"pageNo"`
    PageSize  string `form:"pageSize"`
    Version   string `form:"version"`
    LID       string `form:"lid"`
}

func (r *LiveListOverviewReq) Adjust() {
    r.BuyinAuthReq.Adjust()
}

type ProductListReq struct {
    BuyinAuthReq
    DateType                     string `form:"dateType"`
    BeginDate                    string `form:"beginDate"`
    EndDate                      string `form:"endDate"`
    IsActivity                   bool   `form:"isActivity"`
    ActivityID                   string `form:"activityId"`
    KeyWord                      string `form:"keyWord"`
    IndexSelected                string `form:"indexSelected"`
    SaleType                     int    `form:"saleType"`
    ContentType                  int    `form:"contentType"`
    CateIds                      string `form:"cateIds"`
    CateIdsOriginal              int    `form:"cateIdsOriginal"`
    ProductTab                   int    `form:"productTab"`
    OnlyAbnormal                 bool   `form:"onlyAbnormal"`
    OnlyDropGmv                  bool   `form:"onlyDropGmv"`
    OnlyDropProductShow          bool   `form:"onlyDropProductShow"`
    UseCustomizeGmv              bool   `form:"useCustomizeGmv"`
    UseCustomizeProductShow      bool   `form:"useCustomizeProductShow"`
    AbnormalThresholdGmv         string `form:"abnormalThresholdGmv"`
    AbnormalThresholdProductShow int    `form:"abnormalThresholdProductShow"`
    NewVersion                   bool   `form:"newVersion"`
    PageNo                       int    `form:"pageNo"`
    PageSize                     int    `form:"pageSize"`
    LID                          string `form:"lid"`
}

func (r *ProductListReq) Adjust() {
    if r.PageNo == 0 {
        r.PageNo = 1
    }
    if r.PageSize == 0 {
        r.PageSize = 10
    }
    r.BuyinAuthReq.Adjust()
}

type VideoListReq struct {
    BuyinAuthReq
    PageNo        string `form:"pageNo"`
    PageSize      string `form:"pageSize"`
    BeginDate     string `form:"beginDate"`
    EndDate       string `form:"endDate"`
    DateType      string `form:"dateType"`
    ActivityID    string `form:"activityId"`
    AccountType   string `form:"accountType"`
    AuthorID      string `form:"authorId"`
    RangeType     string `form:"rangeType"`
    CartType      string `form:"cartType"`
    AdType        string `form:"adType"`
    SearchInfo    string `form:"searchInfo"`
    IndexSelected string `form:"indexSelected"`
    SortField     string `form:"sortField"`
    LID           string `form:"lid"`
}

func (r *VideoListReq) Adjust() {
    if r.PageNo == "" {
        r.PageNo = "1"
    }
    if r.PageSize == "" {
        r.PageSize = "10"
    }
    r.BuyinAuthReq.Adjust()
}

type TalentSaleAnalyzeReq struct {
    BuyinAuthReq
    UID   string `form:"uid" binding:"required"`
    Range string `form:"range"`
}

func (r *TalentSaleAnalyzeReq) Adjust() {
    if r.Range == "" {
        r.Range = "30d"
    }
    r.BuyinAuthReq.Adjust()
}