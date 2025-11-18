package buyin

import (
	"encoding/json"
	"fmt"
)

// FansAnalyzeParams 达人粉丝数据分析查询参数
// 与 jx/fans_analyze.go 的查询参数保持一致，使用字符串以便直接传入
// 返回值顺序与 doRequest 保持一致：响应文本、HTTP 状态文本、错误
type FansAnalyzeParams struct {
	EWID      string // ewid参数
	UID       string // uid参数
	FansClub  string // fans_club参数
	WorksType string // works_type参数
}

// FansAnalyze 获取达人粉丝数据分析
func (c *JXClient) FansAnalyze(params *FansAnalyzeParams) (*FansAnalyzeResp, error) {
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"priority":           "u=1, i",
		"referer":            "https://buyin.jinritemai.com/dashboard/servicehall/daren-profile?uid=v2_0a2cb5d3c256be7cf67c20748927e1cd6d48c64a74d43e0a74e2fd428b600e07f12917c55b41a9780894dc5671b51a4b0a3c000000000000000000004fa2a8e2e6ec9b9db570c0c20d68926acaa827ef1482b6bd5896538497959b7ec8ad5d406632130e94501b5d9783edb308e910b2dfff0d18e5ade4c9012001220103d358fd04&enter_from=1&previous_page_name=0%2C5&previous_page_type=0%2C101&search_id=20251025144258D836755603F709DAB2CE&search_source=hand&query=&log_id=20251025144258D836755603F709DAB2CE&module_name=res_source&filter=%E8%A7%86%E9%A2%91%E9%9D%9E%E6%8E%A8%E5%B9%BF%E6%B5%81%E9%87%8F%E5%8D%A0%E6%AF%94%EF%BC%9A%2010%25-30%25&rule_link=https%3A%2F%2Fbuyin.jinritemai.com%2Fdashboard%2Fauthor%2Fconstruct-equity&im_status=99&im_hover_text=%E6%9C%AC%E5%91%A8%E4%BB%85%E6%94%AF%E6%8C%81%E5%90%91%E5%8C%B9%E9%85%8D%E7%B1%BB%E7%9B%AE%E8%BE%BE%E4%BA%BA%E5%8F%91%E9%80%81%E6%B6%88%E6%81%AF%EF%BC%8C%E8%BE%BE%E4%BA%BA%E4%B8%BB%E6%8E%A8%E7%B1%BB%E7%9B%AE%E4%B8%8D%E5%B1%9E%E4%BA%8E%E5%BA%97%E9%93%BA%E7%B1%BB%E7%9B%AE&author_level=6&btm_ppre=a52248.b45827.c0.d0&btm_pre=a52248.b31023.c457762.d68048&btm_show_id=a4871729-e904-4c54-ab98-9c63825b70fe",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         c.config.UserAgent,
	}

	queryParams := map[string]string{
		"ewid":       c.config.EWID,
		"uid":        params.UID,
		"fans_club":  params.FansClub,
		"works_type": params.WorksType,
		"verifyFp":   c.config.VerifyFp,
		"fp":         c.config.Fp,
		"msToken":    c.config.MsToken,
	}

	baseURL := "https://buyin.jinritemai.com/api/authorStatData/authorFansV2"
	resp, _, err := c.doRequest("GET", baseURL, headers, queryParams, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}

	var fansAnalyzeResp FansAnalyzeResp
	_ = json.Unmarshal([]byte(resp), &fansAnalyzeResp)

	if fansAnalyzeResp.Code != 0 {
		return nil, fmt.Errorf("响应错误: %s", fansAnalyzeResp.Msg)
	}

	return &fansAnalyzeResp, nil
}

type FansAnalyzeResp struct {
	St    int    `json:"st"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Extra struct {
		Now   int64  `json:"now"`
		LogId string `json:"log_id"`
	} `json:"extra"`
	Data struct {
		Age struct {
			X       []string `json:"x"`
			Y       []int    `json:"y"`
			Summary []string `json:"summary"`
		} `json:"age"`
		City []struct {
			X       string `json:"x"`
			Y       int    `json:"y"`
			Percent int    `json:"percent"`
		} `json:"city"`
		Device struct {
			X       []string `json:"x"`
			Y       []int    `json:"y"`
			Summary []string `json:"summary"`
		} `json:"device"`
		Gender struct {
			X       []string `json:"x"`
			Y       []int    `json:"y"`
			Summary []string `json:"summary"`
		} `json:"gender"`
		Liveness struct {
			X       []string `json:"x"`
			Y       []int    `json:"y"`
			Summary []string `json:"summary"`
		} `json:"liveness"`
		TrendIncr struct {
			X       []string      `json:"x"`
			Y       []int         `json:"y"`
			Summary []interface{} `json:"summary"`
		} `json:"trend_incr"`
		TrendTotal struct {
			X       []string      `json:"x"`
			Y       []int         `json:"y"`
			Summary []interface{} `json:"summary"`
		} `json:"trend_total"`
		FansVibration        int     `json:"fans_vibration"`
		FansVibrationRatio   float64 `json:"fans_vibration_ratio"`
		StatisticPeriod      string  `json:"statistic_period"`
		ProvinceDistribution []struct {
			X       string `json:"x"`
			Y       int    `json:"y"`
			Percent int    `json:"percent"`
		} `json:"province_distribution"`
		ViewerConclusion string `json:"viewer_conclusion"`
		GmvMainCate      struct {
			X       []string `json:"x"`
			Y       []int    `json:"y"`
			Summary []string `json:"summary"`
		} `json:"gmv_main_cate"`
		AvgPayPrice struct {
			X       []string `json:"x"`
			Y       []int    `json:"y"`
			Summary []string `json:"summary"`
		} `json:"avg_pay_price"`
		RealCity []struct {
			X       string `json:"x"`
			Y       int    `json:"y"`
			Percent int    `json:"percent"`
		} `json:"real_city"`
		Analysis []struct {
			Title string   `json:"title"`
			Empty string   `json:"empty"`
			Value []string `json:"value"`
		} `json:"analysis"`
		CityLevel struct {
			X       []string `json:"x"`
			Y       []int    `json:"y"`
			Summary []string `json:"summary"`
		} `json:"city_level"`
		ConsumerGroup struct {
			X       []string `json:"x"`
			Y       []int    `json:"y"`
			Summary []string `json:"summary"`
		} `json:"consumer_group"`
	} `json:"data"`
}
