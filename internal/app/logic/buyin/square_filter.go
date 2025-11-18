package buyin

import (
	"encoding/json"
	"fmt"
)

// SquareFilterParams 达人广场筛选接口查询参数
// 与 curl 中的查询参数保持一致，统一使用字符串类型
type SquareFilterParams struct {
	Type     string // type 参数
	ReqScene string // req_scene 参数
}

// SquareFilter 获取达人广场筛选条件
func (c *JXClient) SquareFilter(params *SquareFilterParams) (*SquareFilterFile, string, error) {
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

	// 查询参数统一在 doRequest 中进行转义
	queryParams := map[string]string{
		"type":      params.Type,
		"req_scene": params.ReqScene,
		"ewid":      c.config.EWID,
		"verifyFp":  c.config.VerifyFp,
		"fp":        c.config.Fp,
		"msToken":   c.config.MsToken,
	}

	baseURL := "https://buyin.jinritemai.com/square_pc_api/square/filter"
	respText, httpStatus, err := c.doRequest("GET", baseURL, headers, queryParams, nil)
	if err != nil {
		return nil, "", fmt.Errorf("请求失败: %v", err)
	}

	var resp SquareFilterFile
	_ = json.Unmarshal([]byte(respText), &resp)

	return &resp, httpStatus, nil
}

// SquareFilterFile 对应 square_filter.json 的顶层结构
type SquareFilterFile struct {
	Code int              `json:"code"`
	Data SquareFilterData `json:"data"`
}

// SquareFilterData data 字段
type SquareFilterData struct {
	Headers []SquareFilterHeaderGroup `json:"headers"`
}

// SquareFilterHeaderGroup 每个一级分组
type SquareFilterHeaderGroup struct {
	Name   string             `json:"name"`
	Header []SquareFilterItem `json:"header"`
}

// SquareFilterItem 每个筛选项（支持递归 children）
type SquareFilterItem struct {
	Name      string             `json:"name"`
	Type      string             `json:"type"` // e.g. "text" / "cascade"
	FieldName string             `json:"field_name"`
	Value     string             `json:"value"`
	Children  []SquareFilterItem `json:"children,omitempty"`
}
