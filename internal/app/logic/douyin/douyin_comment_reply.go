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
)

// DouyinCommentReply 根据 comment_reply.txt 的 curl 请求实现 GET 评论回复列表接口访问
// 返回值：响应文本、HTTP 状态文本、错误
func (c *DouyinClient) DouyinCommentReply(ctx context.Context, in *types_douyin.CommentReplyListReq) (*types_douyin.CommentData, error) {
	client := &http.Client{}

	out := &types_douyin.CommentData{}

	// 请求头（参考 curl），user-agent 使用客户端传入
	headers := map[string]string{
		"accept":                            "application/json, text/plain, */*",
		"accept-language":                   "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"bd-ticket-guard-client-data":       "eyJ0c19zaWduIjoidHMuMi4zNjE4YzJlZmY1OGNmOWMwZmJmNzg3ZDAwZDUwMjYxZWVmZGQ3MDc4MjE1ZDg1ZmYxYWZhNjc4MWI3OTMwMTc4YzRmYmU4N2QyMzE5Y2YwNTMxODYyNGNlZGExNDkxMWNhNDA2ZGVkYmViZWRkYjJlMzBmY2U4ZDRmYTAyNTc1ZCIsInJlcV9jb250ZW50IjoidGlja2V0LHBhdGgsdGltZXN0YW1wIiwicmVxX3NpZ24iOiJyVFZLY1hqb0lqNmJiMlJPM0xyTTV1WEh0cE1GMFFUNmxHVFBJOFpjbDI0PSIsInRpbWVzdGFtcCI6MTc2MjI0MTY5M30=",
		"bd-ticket-guard-iteration-version": "1",
		"bd-ticket-guard-ree-public-key":    "BJNAg0oNBJTWkwxzsVpUGBhRaPG9iBW/ktG28hIETbpOk9R7lJKywMUEiasQ7qZKvzAY/5N6WCRvKrdQv8hrOVQ=",
		"bd-ticket-guard-version":           "2",
		"bd-ticket-guard-web-sign-type":     "1",
		"bd-ticket-guard-web-version":       "2",
		"origin":                            "https://www.douyin.com",
		"priority":                          "u=1, i",
		"referer":                           "https://www.douyin.com/",
		"sec-ch-ua":                         "\"Chromium\";v=\"142\", \"Microsoft Edge\";v=\"142\", \"Not_A Brand\";v=\"99\"",
		"sec-ch-ua-mobile":                  "?0",
		"sec-ch-ua-platform":                "\"Windows\"",
		"sec-fetch-dest":                    "empty",
		"sec-fetch-mode":                    "cors",
		"sec-fetch-site":                    "same-site",
		"uifid":                             "164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7",
		"user-agent":                        c.UserAgent,
	}

	// 构造查询字符串（顺序与 curl 保持一致，不包含 a_bogus）
	pa := fmt.Sprintf("device_platform=%s&aid=%s&channel=%s&item_id=%s&comment_id=%s&cut_version=%s&cursor=%s&count=%s&item_type=%s&update_version_code=%s&pc_client_type=%s&pc_libra_divert=%s&support_h265=%s&support_dash=%s&cpu_core_num=%s&version_code=%s&version_name=%s&cookie_enabled=%s&screen_width=%s&screen_height=%s&browser_language=%s&browser_platform=%s&browser_name=%s&browser_version=%s&browser_online=%s&engine_name=%s&engine_version=%s&os_name=%s&os_version=%s&device_memory=%s&platform=%s&downlink=%s&effective_type=%s&round_trip_time=%s&webid=%s&uifid=%s&verifyFp=%s&fp=%s&msToken=%s",
		url.QueryEscape("webapp"),
		url.QueryEscape("6383"),
		url.QueryEscape("channel_pc_web"),
		url.QueryEscape(in.ItemId),
		url.QueryEscape(in.CommentId),
		url.QueryEscape("1"),
		url.QueryEscape(in.Cursor),
		url.QueryEscape(in.Count),
		url.QueryEscape("0"),
		url.QueryEscape("170400"),
		url.QueryEscape("1"),
		url.QueryEscape("Windows"),
		url.QueryEscape("0"),
		url.QueryEscape("1"),
		url.QueryEscape("16"),
		url.QueryEscape("170400"),
		url.QueryEscape("17.4.0"),
		url.QueryEscape("true"),
		url.QueryEscape("2560"),
		url.QueryEscape("1080"),
		url.QueryEscape("zh-CN"),
		url.QueryEscape("Win32"),
		url.QueryEscape("Edge"),
		url.QueryEscape("142.0.0.0"),
		url.QueryEscape("true"),
		url.QueryEscape("Blink"),
		url.QueryEscape("142.0.0.0"),
		url.QueryEscape("Windows"),
		url.QueryEscape("10"),
		url.QueryEscape("8"),
		url.QueryEscape("PC"),
		url.QueryEscape("10"),
		url.QueryEscape("4g"),
		url.QueryEscape("50"),
		url.QueryEscape("7513848749303924275"),
		url.QueryEscape("164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7"),
		url.QueryEscape("verify_mgissah8_C6nEvcyk_5pgr_4Qb6_9BGQ_phmZxr733zQ4"),
		url.QueryEscape("verify_mgissah8_C6nEvcyk_5pgr_4Qb6_9BGQ_phmZxr733zQ4"),
		url.QueryEscape("7u8Qg8zTvExZ-xtohdvS8z71rzwM94hXDhnvQ9sEy-PuJdXysxLKZRbmmHHITbNCbjmsMFnJTArS0aqW4j5zkm81TzieKnsqM9GHhud7xZdR2feNJDDq_s2rEO2sxsFrQpiLc5U-phEXlsQG-A9tmaJH124DwbYZUYnylgRTBasLdjn_1taAiQ%3D%3D"),
	)

	// GET 无请求体，da 为空字符串
	da := ""

	// 生成 a_bogus
	aBogus, err := c.generateABogusV2(pa, da)
	if err != nil {
		return out, fmt.Errorf("生成a_bogus失败: %v", err)
	}

	// 构建完整URL（保持与curl一致的主域名）
	baseURL := "https://www-hj.douyin.com/aweme/v1/web/comment/list/reply/"
	fullURL := fmt.Sprintf("%s?%s&a_bogus=%s", baseURL, pa, url.QueryEscape(aBogus))

	// 创建请求
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return out, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// 设置 Cookie 头（来自配置）
	req.Header.Set("Cookie", c.Cookie)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return out, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应，处理可能的gzip压缩
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return out, fmt.Errorf("创建gzip读取器失败: %v", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return out, fmt.Errorf("读取响应失败: %v", err)
	}

	json.Unmarshal(body, &out)
	return out, nil
}
