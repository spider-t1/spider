package buyin

import (
	"encoding/json"
	"fmt"
)

// AllianceProfileParams 达人档案页请求参数
// 参考 alliance_profile.txt 中的 curl，统一使用字符串类型以便直传
type AllianceProfileParams struct {
	UID       string // uid 参数（v2_ 开头的加密UID）
	WorksType string // works_type 参数（内容类型：1等）
}

// AllianceProfileResp 档案页响应（保留 data 为原始 JSON）
type AllianceProfileResp struct {
	Code  int             `json:"code"`
	Data  json.RawMessage `json:"data"`
	LogId string          `json:"log_id"`
	Msg   string          `json:"msg"`
	St    int             `json:"st"`
}

// AllianceProfile 获取达人档案页信息
// GET https://buyin.jinritemai.com/square_pc_api/homePage/author/profile
func (c *JXClient) AllianceProfile(params *AllianceProfileParams) (*AllianceProfileData, error) {
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"priority":           "u=1, i",
		"referer":            "https://buyin.jinritemai.com/dashboard/servicehall/daren-profile",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         c.config.UserAgent,
	}
	if params.WorksType == "" {
		params.WorksType = "1"
	}

	// 查询参数原样传入，由 doRequest 统一进行转义并拼接 a_bogus
	queryParams := map[string]string{
		"ewid":       c.config.EWID,
		"uid":        params.UID,
		"works_type": params.WorksType,
		"verifyFp":   c.config.VerifyFp,
		"fp":         c.config.Fp,
		"msToken":    c.config.MsToken,
	}

	baseURL := "https://buyin.jinritemai.com/square_pc_api/homePage/author/profile"
	respText, _, err := c.doRequest("GET", baseURL, headers, queryParams, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}

	var resp AllianceProfileData
	_ = json.Unmarshal([]byte(respText), &resp)

	if resp.Code != 0 {
		return nil, fmt.Errorf("响应错误: %s", resp.Msg)
	}

	return &resp, nil
}

type AllianceProfileData struct {
	Code int `json:"code"`
	Data struct {
		Level           int      `json:"level"`
		AccountDouyin   string   `json:"account_douyin"`
		Gender          int      `json:"gender"`
		City            string   `json:"city"`
		BindLarkStatus  bool     `json:"bind_lark_status"`
		Nickname        string   `json:"nickname"`
		FansSum         string   `json:"fans_sum"`
		WorksType       []string `json:"works_type"`
		Agency          string   `json:"agency"`
		ProductMainType struct {
			滋补保健 string `json:"滋补保健"`
			美妆   string `json:"美妆"`
		} `json:"product_main_type"`
		ProductMainTypeArray []struct {
			Name string `json:"name"`
			Val  string `json:"val"`
		} `json:"product_main_type_array"`
		Score               string        `json:"score"`
		ReputationLevel     int           `json:"reputation_level"`
		SpecialPrice        int           `json:"special_price"`
		JoinPrice           int           `json:"join_price"`
		Bargaining          int           `json:"bargaining"`
		Duration            int           `json:"duration"`
		InBusiness          int           `json:"in_business"`
		Introduction        string        `json:"introduction"`
		SaleType            string        `json:"sale_type"`
		RecReason           []interface{} `json:"rec_reason"`
		DarenPlazaRecReason struct {
			Tag   int    `json:"tag"`
			Value string `json:"value"`
		} `json:"daren_plaza_rec_reason"`
		Avatar           string        `json:"avatar"`
		ErrorMsg         string        `json:"error_msg"`
		ShareUrlDouyin   string        `json:"share_url_douyin"`
		RecommendReasons []interface{} `json:"recommend_reasons"`
		CreditScore      int           `json:"credit_score"`
		IntentionCatgory []interface{} `json:"intention_catgory"`
		CooperateMode    string        `json:"cooperate_mode"`
		CommissionRatio  string        `json:"commission_ratio"`
		SellRequirement  struct {
			ChargeStandard struct {
				Tag   int    `json:"tag"`
				Value string `json:"value"`
			} `json:"charge_standard"`
			Requirement int `json:"requirement"`
		} `json:"sell_requirement"`
		DarkHorses              []interface{} `json:"dark_horses"`
		HighOnlineReplyRate     bool          `json:"high_online_reply_rate"`
		HighInvitationReplyRate bool          `json:"high_invitation_reply_rate"`
		InsitutionId            string        `json:"insitution_id"`
		WebHomepageUrl          string        `json:"web_homepage_url"`
		HighCooperation         bool          `json:"high_cooperation"`
		IsStar                  bool          `json:"is_star"`
		Tags                    []interface{} `json:"tags"`
		ActInfo                 interface{}   `json:"act_info"`
		AuthorRecReasons        []struct {
			TagType  int    `json:"tag_type"`
			Reason   string `json:"reason"`
			Extra    string `json:"extra"`
			ExtraMap struct {
				Category string `json:"category"`
				Tips     string `json:"tips"`
			} `json:"extra_map"`
			LogParams struct {
				Category string `json:"category"`
				Reason   string `json:"reason"`
				TagType  string `json:"tag_type"`
				Tips     string `json:"tips"`
			} `json:"log_params"`
		} `json:"author_rec_reasons"`
	} `json:"data"`
	LogId string `json:"log_id"`
	Msg   string `json:"msg"`
	St    int    `json:"st"`
}
