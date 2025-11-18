package buyin

// LiveListOverviewParams 直播列表概览参数
type LiveListOverviewParams struct {
	AType     string // a_type参数
	DateType  string // date_type参数
	BeginDate string // begin_date参数
	EndDate   string // end_date参数
	SortField string // sort_field参数
	IsAsc     string // is_asc参数
	PageNo    string // page_no参数
	PageSize  string // page_size参数
	Version   string // version参数
	LID       string // _lid参数
}

// LiveListOverview 获取直播列表概览
func (c *JXClient) LiveListOverview(params LiveListOverviewParams) (string, string, error) {
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"agw-js-conv":        "str",
		"priority":           "u=1, i",
		"referer":            "https://compass.jinritemai.com/shop/live-list?from=sy&from_page=%2Fshop&btm_ppre=a0.b0.c0.d0&btm_pre=a6187.b01487.c0.d0&btm_show_id=0c524c6d-a267-4348-8363-b0cb9aeee683",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         c.config.UserAgent,
	}

	queryParams := map[string]string{
		"a_type":     params.AType,
		"date_type":  params.DateType,
		"begin_date": params.BeginDate,
		"end_date":   params.EndDate,
		"sort_field": params.SortField,
		"is_asc":     params.IsAsc,
		"page_no":    params.PageNo,
		"page_size":  params.PageSize,
		"version":    params.Version,
		"_lid":       params.LID,
		"verifyFp":   c.config.VerifyFp,
		"fp":         c.config.Fp,
		"msToken":    c.config.MsToken,
	}

	baseURL := "https://compass.jinritemai.com/compass_api/shop/live/live_list/overview"
	return c.doRequest("GET", baseURL, headers, queryParams, nil)
}
