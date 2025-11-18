package buyin

// GetFrontPageOperatingData 获取首页运营数据
func (c *JXClient) GetFrontPageOperatingData(timeRange string) (string, string, error) {
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"priority":           "u=1, i",
		"referer":            "https://buyin.jinritemai.com/dashboard",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         c.config.UserAgent,
	}

	queryParams := map[string]string{
		"time_range": timeRange,
		"ewid":       c.config.EWID,
		"verifyFp":   c.config.VerifyFp,
		"fp":         c.config.Fp,
		"msToken":    c.config.MsToken,
	}

	baseURL := "https://buyin.jinritemai.com/api/frontpage/getFrontPageOperatingData"
	return c.doRequest("GET", baseURL, headers, queryParams, nil)
}
