package buyin

import (
	"fmt"
)

// ProductListRequest 产品列表请求参数
type ProductListRequest struct {
	DateType                     string `json:"date_type"`
	BeginDate                    string `json:"begin_date"`
	EndDate                      string `json:"end_date"`
	IsActivity                   bool   `json:"is_activity"`
	ActivityID                   string `json:"activity_id"`
	KeyWord                      string `json:"key_word"`
	IndexSelected                string `json:"index_selected"`
	SaleType                     int    `json:"sale_type"`
	ContentType                  int    `json:"content_type"`
	CateIds                      string `json:"cate_ids"`
	CateIdsOriginal              int    `json:"cate_ids_original"`
	ProductTab                   int    `json:"product_tab"`
	OnlyAbnormal                 bool   `json:"only_abnormal"`
	OnlyDropGmv                  bool   `json:"only_drop_gmv"`
	OnlyDropProductShow          bool   `json:"only_drop_product_show"`
	UseCustomizeGmv              bool   `json:"use_customize_gmv"`
	UseCustomizeProductShow      bool   `json:"use_customize_product_show"`
	AbnormalThresholdGmv         string `json:"abnormal_threshold_gmv"`
	AbnormalThresholdProductShow int    `json:"abnormal_threshold_product_show"`
	NewVersion                   bool   `json:"new_version"`
	PageNo                       int    `json:"page_no"`
	PageSize                     int    `json:"page_size"`
	LID                          string `json:"_lid"`
}

// GetProductList 获取产品列表
func (c *JXClient) GetProductList(request ProductListRequest) (string, string, error) {
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"agw-js-conv":        "str",
		"priority":           "u=1, i",
		"referer":            "https://compass.jinritemai.com/shop/commodity/product-list?from_page=%2Fshop%2Flive-detail&from=sy&btm_ppre=a6187.b4991.c73335.d015213_i3&btm_pre=a6187.b6554.c0.d0&btm_show_id=e5b350a6-103b-46eb-9914-60203feea1cf",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0",
	}

	// 构建查询参数（值原样传入，统一在 doRequest 中转义）
	queryParams := map[string]string{
		"date_type":                       request.DateType,
		"begin_date":                      request.BeginDate,
		"end_date":                        request.EndDate,
		"is_activity":                     fmt.Sprintf("%t", request.IsActivity),
		"activity_id":                     request.ActivityID,
		"key_word":                        request.KeyWord,
		"index_selected":                  request.IndexSelected,
		"sale_type":                       fmt.Sprintf("%d", request.SaleType),
		"content_type":                    fmt.Sprintf("%d", request.ContentType),
		"cate_ids":                        request.CateIds,
		"cate_ids_original":               fmt.Sprintf("%d", request.CateIdsOriginal),
		"product_tab":                     fmt.Sprintf("%d", request.ProductTab),
		"only_abnormal":                   fmt.Sprintf("%t", request.OnlyAbnormal),
		"only_drop_gmv":                   fmt.Sprintf("%t", request.OnlyDropGmv),
		"only_drop_product_show":          fmt.Sprintf("%t", request.OnlyDropProductShow),
		"use_customize_gmv":               fmt.Sprintf("%t", request.UseCustomizeGmv),
		"use_customize_product_show":      fmt.Sprintf("%t", request.UseCustomizeProductShow),
		"abnormal_threshold_gmv":          request.AbnormalThresholdGmv,
		"abnormal_threshold_product_show": fmt.Sprintf("%d", request.AbnormalThresholdProductShow),
		"new_version":                     fmt.Sprintf("%t", request.NewVersion),
		"page_no":                         fmt.Sprintf("%d", request.PageNo),
		"page_size":                       fmt.Sprintf("%d", request.PageSize),
		"_lid":                            request.LID,
		"verifyFp":                        c.config.VerifyFp,
		"fp":                              c.config.Fp,
		"msToken":                         c.config.MsToken,
	}

	baseURL := "https://compass.jinritemai.com/compass_api/shop/product/product/product_list"
	res, status, err := c.doRequest("GET", baseURL, headers, queryParams, nil)
	if err != nil {
		return "", "", err
	}
	return status, res, nil
}
