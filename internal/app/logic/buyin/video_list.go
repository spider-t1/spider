package buyin

// VideoListParams 视频列表查询参数
// 与 jx/video_list.go 的查询参数保持一致，使用字符串以便直接传入
// 返回值顺序与 doRequest 保持一致：响应文本、HTTP 状态文本、错误
type VideoListParams struct {
	PageNo        string
	PageSize      string
	BeginDate     string
	EndDate       string
	DateType      string
	ActivityID    string
	AccountType   string
	AuthorID      string
	RangeType     string
	CartType      string // 1 挂车 2非挂车
	AdType        string
	SearchInfo    string
	IndexSelected string
	SortField     string
	LID           string // _lid
}

// VideoList 获取视频列表
func (c *JXClient) VideoList(params VideoListParams) (string, string, error) {
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"agw-js-conv":        "str",
		"priority":           "u=1, i",
		"referer":            "https://compass.jinritemai.com/shop/video/cooperate?from_page=%2Fshop%2Fvideo%2Foverview&btm_ppre=a6187.b904798.c0.d0&btm_pre=a6187.b916250.c0.d0&btm_show_id=998b10e9-7910-47c7-9564-8c30a2ca4b65",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         c.config.UserAgent,
	}

	queryParams := map[string]string{
		"page_no":        params.PageNo,
		"page_size":      params.PageSize,
		"begin_date":     params.BeginDate,
		"end_date":       params.EndDate,
		"date_type":      params.DateType,
		"activity_id":    params.ActivityID,
		"account_type":   params.AccountType,
		"author_id":      params.AuthorID,
		"range_type":     params.RangeType,
		"cart_type":      params.CartType,
		"ad_type":        params.AdType,
		"search_info":    params.SearchInfo,
		"index_selected": params.IndexSelected,
		"sort_field":     params.SortField,
		"_lid":           params.LID,
		"verifyFp":       c.config.VerifyFp,
		"fp":             c.config.Fp,
		"msToken":        c.config.MsToken,
	}

	baseURL := "https://compass.jinritemai.com/compass_api/shop/video/overview/video_list"
	return c.doRequest("GET", baseURL, headers, queryParams, nil)
}
