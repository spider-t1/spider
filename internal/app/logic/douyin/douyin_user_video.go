package douyin

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"spider/internal/app/types/types_douyin"
	"spider/pkg/logger"
)

// DouyinUserVideo 根据 douyin_user_video.txt 的 curl 请求实现 GET 接口访问
// 返回值：响应文本、HTTP 状态文本、错误
func (c *DouyinClient) DouyinUserVideo(ctx context.Context, secId string) (*types_douyin.DyOneVideoInfo, error) {
	client := &http.Client{}

	dyVideo := types_douyin.DyOneVideoInfo{}

	logger.InfoWithContext(ctx, fmt.Sprintf("请求抖音用户视频列表: secId %s", secId))

	// 请求头（参考 curl）
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"origin":             "https://www.douyin.com",
		"priority":           "u=1, i",
		"referer":            "https://www.douyin.com/",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-site",
		"uifid":              "164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0",
	}

	// 将 curl 的 cookie 作为整串放入 Header

	// 构造查询字符串（顺序与 curl 保持一致，不包含 a_bogus）
	pa := fmt.Sprintf("device_platform=%s&aid=%s&channel=%s&sec_user_id=%s&max_cursor=%s&locate_query=%s&show_live_replay_strategy=%s&need_time_list=%s&time_list_query=%s&whale_cut_token=%s&cut_version=%s&count=%s&publish_video_strategy_type=%s&from_user_page=%s&update_version_code=%s&pc_client_type=%s&pc_libra_divert=%s&support_h265=%s&support_dash=%s&cpu_core_num=%s&version_code=%s&version_name=%s&cookie_enabled=%s&screen_width=%s&screen_height=%s&browser_language=%s&browser_platform=%s&browser_name=%s&browser_version=%s&browser_online=%s&engine_name=%s&engine_version=%s&os_name=%s&os_version=%s&device_memory=%s&platform=%s&downlink=%s&effective_type=%s&round_trip_time=%s&webid=%s&uifid=%s&msToken=%s&verifyFp=%s&fp=%s",
		url.QueryEscape("webapp"),
		url.QueryEscape("6383"),
		url.QueryEscape("channel_pc_web"),
		url.QueryEscape(secId),
		url.QueryEscape("0"),
		url.QueryEscape("false"),
		url.QueryEscape("1"),
		url.QueryEscape("1"),
		url.QueryEscape("0"),
		url.QueryEscape(""),
		url.QueryEscape("1"),
		url.QueryEscape("18"),
		url.QueryEscape("2"),
		url.QueryEscape("1"),
		url.QueryEscape("170400"),
		url.QueryEscape("1"),
		url.QueryEscape("Windows"),
		url.QueryEscape("0"),
		url.QueryEscape("1"),
		url.QueryEscape("16"),
		url.QueryEscape("290100"),
		url.QueryEscape("29.1.0"),
		url.QueryEscape("true"),
		url.QueryEscape("2560"),
		url.QueryEscape("1080"),
		url.QueryEscape("zh-CN"),
		url.QueryEscape("Win32"),
		url.QueryEscape("Edge"),
		url.QueryEscape("141.0.0.0"),
		url.QueryEscape("true"),
		url.QueryEscape("Blink"),
		url.QueryEscape("141.0.0.0"),
		url.QueryEscape("Windows"),
		url.QueryEscape("10"),
		url.QueryEscape("8"),
		url.QueryEscape("PC"),
		url.QueryEscape("10"),
		url.QueryEscape("4g"),
		url.QueryEscape("0"),
		url.QueryEscape("7513848749303924275"),
		url.QueryEscape("164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7"),
		url.QueryEscape("wbuZiuasIYn0nD_aF8MHSrp6RjarEoIOsYtvSRIebW6C2vu5X2WT9y_ESK1Qitpf_oTtOIG4AtikOxW3lZWD16VHRjVcjK2lG1hb7EUqG6fPVGm-pmcXYpbHUJLy2TfSzw5R1RIwiOy7srDujN4dLz-J-f6D4pJ0IwZG8hWVirWtEhqpiLzWbg%3D%3D"),
		url.QueryEscape("verify_mgissah8_C6nEvcyk_5pgr_4Qb6_9BGQ_phmZxr733zQ4"),
		url.QueryEscape("verify_mgissah8_C6nEvcyk_5pgr_4Qb6_9BGQ_phmZxr733zQ4"),
	)

	// GET 无请求体，da 为空字符串
	da := ""

	// 生成 a_bogus
	aBogus, err := c.generateABogusV2(pa, da)
	if err != nil {
		return &dyVideo, fmt.Errorf("生成a_bogus失败: %v", err)
	}

	logger.InfoWithContext(ctx, fmt.Sprintf("请求参数: %s", pa))

	// 构建完整URL
	baseURL := "https://www-hj.douyin.com/aweme/v1/web/aweme/post/"
	fullURL := fmt.Sprintf("%s?%s&a_bogus=%s", baseURL, pa, url.QueryEscape(aBogus))

	logger.InfoWithContext(ctx, fmt.Sprintf("请求URL: %s", fullURL))
	// 创建请求
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return &dyVideo, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// 设置 Cookie 头
	req.Header.Set("Cookie", c.Cookie)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return &dyVideo, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应，处理可能的gzip压缩
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return &dyVideo, fmt.Errorf("创建gzip读取器失败: %v", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	body, err := io.ReadAll(reader)
	json.Unmarshal(body, &dyVideo)

	logger.InfoWithContext(ctx, fmt.Sprintf("响应数据: %s", string(body)))
	if err != nil {
		return &dyVideo, fmt.Errorf("读取响应失败: %v", err)
	}

	return &dyVideo, nil
}

func (c *DouyinClient) generateABogusV2(pa, da string) (string, error) {
	// 直接调用Go版本的Enc函数
	result := Enc(pa, da, c.UserAgent)
	return result, nil
}
