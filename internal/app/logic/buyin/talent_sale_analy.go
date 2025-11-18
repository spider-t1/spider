package buyin

import (
	"encoding/json"
)

// TalentSaleAnalyzeParams 达人销售分析查询参数
type TalentSaleAnalyzeParams struct {
	UID   string `json:"uid"`   // uid参数
	Range string `json:"range"` // 时间范围参数，如"30d"
}

type TalentSaleAnalyzeResp struct {
	St    int    `json:"st"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Extra struct {
		Now   int64  `json:"now"`
		LogId string `json:"log_id"`
	} `json:"extra"`
	Data struct {
		KeyData struct {
			SaleLow          int `json:"sale_low"`
			SaleHigh         int `json:"sale_high"`
			PromotionSum     int `json:"promotion_sum"`
			CooperateShopNum int `json:"cooperate_shop_num"`
			AveragePrice     int `json:"average_price"`
			SaleValue        int `json:"sale_value"`
			SaleStatus       int `json:"sale_status"`
		} `json:"key_data"`
		LiveData struct {
			Percentage        int    `json:"percentage"`
			SaleLow           int    `json:"sale_low"`
			SaleHigh          int    `json:"sale_high"`
			GPMLow            int    `json:"GPM_low"`
			GPMHigh           int    `json:"GPM_high"`
			PromotionSum      int    `json:"promotion_sum"`
			CooperateShopNum  int    `json:"cooperate_shop_num"`
			AveragePrice      int    `json:"average_price"`
			SaleValue         int    `json:"sale_value"`
			GPMValue          int    `json:"GPM_value"`
			SaleStatus        int    `json:"sale_status"`
			GPMStatus         int    `json:"GPM_status"`
			RecommendRate     string `json:"recommend_rate"`
			HighRecommendRate bool   `json:"high_recommend_rate"`
		} `json:"live_data"`
		VideoData struct {
			Percentage        int    `json:"percentage"`
			SaleLow           int    `json:"sale_low"`
			SaleHigh          int    `json:"sale_high"`
			GPMLow            int    `json:"GPM_low"`
			GPMHigh           int    `json:"GPM_high"`
			PromotionSum      int    `json:"promotion_sum"`
			CooperateShopNum  int    `json:"cooperate_shop_num"`
			AveragePrice      int    `json:"average_price"`
			SaleValue         int    `json:"sale_value"`
			GPMValue          int    `json:"GPM_value"`
			SaleStatus        int    `json:"sale_status"`
			GPMStatus         int    `json:"GPM_status"`
			RecommendRate     string `json:"recommend_rate"`
			HighRecommendRate bool   `json:"high_recommend_rate"`
		} `json:"video_data"`
		DateSection []string `json:"date_section"`
	} `json:"data"`
}

// TalentSaleAnaly 达人销售分析接口
// 返回值：响应文本、HTTP 状态文本、错误
func (c *JXClient) TalentSaleAnaly(params TalentSaleAnalyzeParams) (*TalentSaleAnalyzeResp, string, error) {

	if params.Range == "" {
		params.Range = "30d"
	}

	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"priority":           "u=1, i",
		"referer":            "https://buyin.jinritemai.com/dashboard/servicehall/daren-profile?uid=v2_0a27bf94e321655b7761f869ced6f9d1b50ee85618758deee6eba246d02f86f0e190b1d198728693e81a4b0a3c000000000000000000004fa12d2530d846f60de7c0f982aca22525a7ea4cd66f9a4a3c5b959543a865778153d888a1fe3d29f48662c8681cce5f5b9b1085d4ff0d18e5ade4c90120012201036d2c3fcc&enter_from=1&previous_page_name=0%2C5&previous_page_type=0%2C100&search_id=&search_source=&query=&log_id=20251024105407E059EBFE8C05F64F7189&module_name=recommend&filter=&rule_link=https%3A%2F%2Fbuyin.jinritemai.com%2Fdashboard%2Fauthor%2Fconstruct-equity&im_status=1&im_hover_text=%E6%9C%AC%E5%91%A8%E5%90%91%20LV2-LV3%20%E7%9A%84%E8%BE%BE%E4%BA%BA%E5%8F%91%E8%B5%B7%E9%82%80%E7%BA%A6%E5%B8%A6%E8%B4%A7%E6%AC%A1%E6%95%B0%E4%B8%8A%E9%99%90%E4%B8%BA%20203%2F1100%EF%BC%8C%E5%90%84%E7%AD%89%E7%BA%A7%E8%BE%BE%E4%BA%BA%E5%BB%BA%E8%81%94%E4%B8%8A%E9%99%90%E5%8F%AF&author_level=2&btm_ppre=a52248.b45827.c0.d0&btm_pre=a52248.b31023.c0.d0&btm_show_id=a4871729-e904-4c54-ab98-9c63825b70fe",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         c.config.UserAgent,
	}

	// 构建查询参数
	queryParams := map[string]string{
		"ewid":     c.config.EWID,
		"uid":      params.UID,
		"range":    params.Range,
		"verifyFp": c.config.VerifyFp,
		"fp":       c.config.Fp,
		"msToken":  c.config.MsToken,
	}

	// 如果参数结构体中的某些字段为空，使用客户端配置中的默认值
	if queryParams["verifyFp"] == "" {
		queryParams["verifyFp"] = c.config.VerifyFp
	}
	if queryParams["fp"] == "" {
		queryParams["fp"] = c.config.Fp
	}
	if queryParams["msToken"] == "" {
		queryParams["msToken"] = c.config.MsToken
	}

	baseURL := "https://buyin.jinritemai.com/api/authorStatData/salesAnalyseV2"
	request, s, err := c.doRequest("GET", baseURL, headers, queryParams, nil)
	var resp TalentSaleAnalyzeResp
	json.Unmarshal([]byte(request), &resp)
	return &resp, s, err
}
