package buyin

import (
	"encoding/json"
	"fmt"
)

// GetShopUserInfo 获取当前精选联盟账号的用户信息
// 参考 curl: https://buyin.jinritemai.com/index/getUser
// GET 参数：verifyFp、fp、msToken（a_bogus 由 doRequest 自动生成并追加）
func (c *JXClient) GetShopUserInfo() (*ShopUserInfoResp, error) {
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"priority":           "u=1, i",
		"referer":            "https://buyin.jinritemai.com/dashboard/servicehall/daren-square",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         c.config.UserAgent,
	}

	// 按照 curl 仅携带 verifyFp、fp、msToken 参数（a_bogus 在 doRequest 内生成并追加）
	queryParams := map[string]string{
		"verifyFp": c.config.VerifyFp,
		"fp":       c.config.Fp,
		"msToken":  c.config.MsToken,
	}

	baseURL := "https://buyin.jinritemai.com/index/getUser"
	respText, _, err := c.doRequest("GET", baseURL, headers, queryParams, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}

	var resp ShopUserInfoResp
	_ = json.Unmarshal([]byte(respText), &resp)
	if resp.Code != 0 {
		return nil, fmt.Errorf("响应错误: %s", resp.Msg)
	}
	if resp.Data.ShopName == "" {
		return nil, fmt.Errorf("响应错误: %s", resp.Msg)
	}
	return &resp, nil
}

// ShopUserInfoResp 店铺用户信息响应结构（保留 data 为原始 JSON）
type ShopUserInfoResp struct {
	Code int `json:"code"`
	Data struct {
		ShopName         string `json:"shop_name"`
		Status           int    `json:"status"`
		CheckLiveBase    bool   `json:"check_live_base"`
		MarginHints      string `json:"margin_hints"`
		LoginUserId      string `json:"login_user_id"`
		ShopType         int    `json:"shop_type"`
		UserRole         int    `json:"user_role"`
		AccountAvatar    string `json:"account_avatar"`
		HasBindStar      int    `json:"has_bind_star"`
		AccountStatus    int    `json:"account_status"`
		IsBanned         bool   `json:"is_banned"`
		ConnectionConfig struct {
			UpPersistTypes []interface{} `json:"up_persist_types"`
			PersistTypes   []int         `json:"persist_types"`
			TtWid          string        `json:"tt_wid"`
			AccessKey      string        `json:"access_key"`
		} `json:"connection_config"`
		ShopNeedAgreePro   bool   `json:"shop_need_agree_pro"`
		UserId             string `json:"user_id"`
		ShopId             string `json:"shop_id"`
		UserApp            int    `json:"user_app"`
		ShopTypeChild      int    `json:"shop_type_child"`
		AccountChildType   int    `json:"account_child_type"`
		AgreePro           int    `json:"agree_pro"`
		AccountChildStatus int    `json:"account_child_status"`
		OriginUid          string `json:"origin_uid"`
		UserName           string `json:"user_name"`
		CheckStatus        int    `json:"check_status"`
		DoudianShopId      int    `json:"doudian_shop_id"`
		IsVsgShop          bool   `json:"is_vsg_shop"`
	} `json:"data"`
	LogId string `json:"log_id"`
	Msg   string `json:"msg"`
	St    int    `json:"st"`
}
