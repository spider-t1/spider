package douyin

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"spider/internal/app/types/types_douyin"
	"spider/pkg/logger"
)

// DouyinUserSearch 抖音用户搜索接口
// 返回值：响应文本、HTTP 状态文本、错误
func (c *DouyinClient) DouyinUserSearch(ctx context.Context, keyword string) (*types_douyin.DouyinSearchResp, error) {
	client := &http.Client{}
	var dy types_douyin.DouyinSearchResp

	logger.InfoWithContext(ctx, fmt.Sprintf("DouyinUserSearch 请求抖音用户搜索: keyword %s", keyword))

	// 请求头（参考 curl）
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"priority":           "u=1, i",
		"referer":            "https://www.douyin.com/user/MS4wLjABAAAAyiUBvOEjnzVYHl3xyIopxDpk1e7ECR/search/qiufengchuilai?type=user",
		"sec-ch-ua":          "\"Microsoft Edge\";v=\"141\", \"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"141\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0",
		"uifid":              "164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7",
	}

	// 将 curl 的 cookie 作为整串放入 Header
	//cookieStr := `UIFID_TEMP=164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45441db5afa96ff09f5b0170e1b31058eb70a505d3baced3e44b6de3c02f2936f84; WebUgChannelId=%2230001%22; xgplayer_user_id=888012742023; fpk1=U2FsdGVkX19BOFXQkOaii8ROHXtRy7MVzt54lbT0LK2taCL1QvG/jGZPoY4xzD0hihxUPoRHPzwdfeH8boxd9A==; fpk2=800683566637788f812c9cb58711ba4c; bd_ticket_guard_client_web_domain=2; UIFID=164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7; SearchMultiColumnLandingAbVer=1; SEARCH_RESULT_LIST_TYPE=%22multi%22; n_mh=Hn11pbjmqhtFQ6NbHB6BtvAxooZl10kP99YFV9cRZ8g; SelfTabRedDotControl=%5B%5D; my_rd=2; live_use_vvc=%22false%22; xgplayer_device_id=61847540687; enter_pc_once=1; d_ticket=9151353157e242473258235b66749b47d9a72; passport_assist_user=CkERozBE1KjJDCm6x2WlabcTiwnbR_tcgfuzU6kQQYe0LdCgA8GUFQQP8cvQkgh1cYK2OTENErT6aso0wiV5TH_Z0BpKCjwAAAAAAAAAAAAAT1HMv2v-0GEWxxM2txtqIiClBsZvR3gtWGiNkYTJcF4qMsFK3Kg20fVfR9_rVAm0s9QQ5dL4DRiJr9ZUIAEiAQNDpFIw; uid_tt=ef01891409b3e784d797278b4b780478; uid_tt_ss=ef01891409b3e784d797278b4b780478; sid_tt=de65f702d51d4b162ada81667eb342ad; sessionid=de65f702d51d4b162ada81667eb342ad; sessionid_ss=de65f702d51d4b162ada81667eb342ad; is_staff_user=false; login_time=1754389635201; __security_mc_1_s_sdk_cert_key=7fd9ce93-4989-bac0; __live_version__=%221.1.4.472%22; __security_mc_1_s_sdk_crypt_sdk=59402a63-4c4c-bd7f; passport_csrf_token=7f744c92495d5355cd4db2fd794e03bd; passport_csrf_token_default=7f744c92495d5355cd4db2fd794e03bd; dy_swidth=2560; dy_sheight=1080; s_v_web_id=verify_mgissah8_C6nEvcyk_5pgr_4Qb6_9BGQ_phmZxr733zQ4; __security_mc_1_s_sdk_sign_data_key_web_protect=65abbc52-4825-b33a; is_dash_user=1; volume_info=%7B%22isUserMute%22%3Afalse%2C%22isMute%22%3Atrue%2C%22volume%22%3A0.992%7D; sid_guard=de65f702d51d4b162ada81667eb342ad%7C1760593524%7C5184000%7CMon%2C+15-Dec-2025+05%3A45%3A24+GMT; sid_ucp_v1=1.0.0-KGIzZjc2MmE2ZjUzNzFhN2Y5ODdjODY3NGY2ZjIxZDJlNDFkNjg2MWQKIQjDoaHu5_TOAxD0jMLHBhjvMSAMMIq9n-sFOAJA8QdIBBoCaGwiIGRlNjVmNzAyZDUxZDRiMTYyYWRhODE2NjdlYjM0MmFk; ssid_ucp_v1=1.0.0-KGIzZjc2MmE2ZjUzNzFhN2Y5ODdjODY3NGY2ZjIxZDJlNDFkNjg2MWQKIQjDoaHu5_TOAxD0jMLHBhjvMSAMMIq9n-sFOAJA8QdIBBoCaGwiIGRlNjVmNzAyZDUxZDRiMTYyYWRhODE2NjdlYjM0MmFk; session_tlb_tag=sttt%7C5%7C3mX3AtUdSxYq2oFmfrNCrf________-ncq7HkaE6S1AEQwfWkJK7r67v0JN2ODxFwuBqPm8UFmA%3D; download_guide=%223%2F20251021%2F0%22; stream_player_status_params=%22%7B%5C%22is_auto_play%5C%22%3A0%2C%5C%22is_full_screen%5C%22%3A0%2C%5C%22is_full_webscreen%5C%22%3A0%2C%5C%22is_mute%5C%22%3A1%2C%5C%22is_speed%5C%22%3A1%2C%5C%22is_visible%5C%22%3A0%7D%22; douyin.com; xg_device_score=7.90435294117647; device_web_cpu_core=16; device_web_memory_size=8; architecture=amd64; strategyABtestKey=%221761210131.606%22; ttwid=1%7CCJxcNR-0EqvaaIyyL9mz09G6TkkSY1nersicZtmJqTY%7C1761210131%7C4671d3ef036050066f7a4994723c8f718e46bf8b941f4b10e22fb867b0a403b1; publish_badge_show_info=%220%2C0%2C0%2C1761210135476%22; __ac_nonce=068f9f92a003208572322; __ac_signature=_02B4Z6wo00f01jrV5kQAAIDDQDwSoj1WSn469eLAAOZaff; gfkadpd=1243,16718; FOLLOW_NUMBER_YELLOW_POINT_INFO=%22MS4wLjABAAAAdvkpdIfz1CAW9LFsCWeavAZa_vrx108_2f6Z817LVOI61aYpO-esUKpoIUWnbFPw%2F1761235200000%2F0%2F0%2F1761214179080%22; playRecommendGuideTagCount=9; totalRecommendGuideTagCount=9; gulu_source_res=eyJwX2luIjoiZjI1NzFkMzg0MDZkYWFhM2I1MGFkY2E0MjgxMDI4N2VmMDEwMDcxYjQzNTA2ZWJkY2RlOGYxZDZmMjYyZWQ0NCJ9; stream_recommend_feed_params=%22%7B%5C%22cookie_enabled%5C%22%3Atrue%2C%5C%22screen_width%5C%22%3A2560%2C%5C%22screen_height%5C%22%3A1080%2C%5C%22browser_online%5C%22%3Atrue%2C%5C%22cpu_core_num%5C%22%3A16%2C%5C%22device_memory%5C%22%3A8%2C%5C%22downlink%5C%22%3A10%2C%5C%22effective_type%5C%22%3A%5C%224g%5C%22%2C%5C%22round_trip_time%5C%22%3A0%7D%22; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtcmVlLXB1YmxpYy1rZXkiOiJCSk5BZzBvTkJKVFdrd3h6c1ZwVUdCaFJhUEc5aUJXL2t0RzI4aElFVGJwT2s5UjdsSkt5d01VRWlhc1E3cVpLdnpBWS81TjZXQ1J2S3JkUXY4aHJPVlE9IiwiYmQtdGlja2V0LWd1YXJkLXdlYi12ZXJzaW9uIjoyfQ%3D%3D; home_can_add_dy_2_desktop=%221%22; odin_tt=34dd63395effea432d4a55c3f588d31c5b0dd88c866e802e503a50fcf5520a61e55d2a10d8abca0ee1375a8b80b6f64e440a70ed82e88120bb979a4684fd206d; bd_ticket_guard_client_data_v2=eyJyZWVfcHVibGljX2tleSI6IkJKTkFnMG9OQkpUV2t3eHpzVnBVR0JoUmFQRzlpQlcva3RHMjhoSUVUYnBPazlSN2xKS3l3TVVFaWFzUTdxWkt2ekFZLzVONldDUnZLcmRRdjhock9WUT0iLCJ0c19zaWduIjoidHMuMi4zNjE4YzJlZmY1OGNmOWMwZmJmNzg3ZDAwZDUwMjYxZWVmZGQ3MDc4MjE1ZDg1ZmYxYWZhNjc4MWI3OTMwMTc4YzRmYmU4N2QyMzE5Y2YwNTMxODYyNGNlZGExNDkxMWNhNDA2ZGVkYmViZWRkYjJlMzBmY2U4ZDRmYTAyNTc1ZCIsInJlcV9jb250ZW50Ijoic2VjX3RzIiwicmVxX3NpZ24iOiI5R0xFTnBJcG50V1Evc1cyanJNd2hvTzZ3dEMwVXAyN1cxK0RqSkhqZWFzPSIsInNlY190cyI6IiNxMUh0WjlVSU9TYk9zcnJieGFoaVF1UU55d3ZDOGh5d29SUTRFem0rd3NNWEg4T0xtNDVOZXBDU2I5TDcifQ%3D%3D; biz_trace_id=ffa7e2f1; sdk_source_info=7e276470716a68645a606960273f276364697660272927676c715a6d6069756077273f276364697660272927666d776a68605a607d71606b766c6a6b5a7666776c7571273f275e58272927666a6b766a69605a696c6061273f27636469766027292762696a6764695a7364776c6467696076273f275e582729277672715a646971273f2763646976602729277f6b5a666475273f2763646976602729276d6a6e5a6b6a716c273f2763646976602729276c6b6f5a7f6367273f27636469766027292771273f2737323535353d36343734333234272927676c715a75776a716a666a69273f2763646976602778; bit_env=BGLFkfHpeE7n-h2km7L-AqTkcG0gXvxIXJO5lJkXgcIuDsXX4zIxGt8DIe8HOgGwqG8KhqjXIl1f-02FuhfYxHja2HZ1nd311q0d49ZmDTBx1ek4LvZCnYTAJwu-Ak76bhWu2_MobmatTxdnQztGBUCdxWoyZIVeOffnZRaeesp3giKERDKyExwY3oRDNlLSIXS9jhvbKjPecWZYDAPBdp6VJwqDCnI_fhjWVNEletrGX54njLE0A7ukHrDb5u2ovmRru-hq6N5THfD_WhZjB5hKEkmqVCkSbTzbZgA-Mj9-xH3_hy9y0FPi9pk2mEFhYch8PiZsUmq_q0JOdQONmysLanHu-xaG_Om5Lvmu3AU81jnmsNQDs8NFKHNXWeHPbZq7JfqjNs2f8lADPeClRHymLhHXIarL60xAJ0Mz5uz40hwpn3m2t5SUrmpHc4s9Av3-qANKygMa9j8UrSMI_8ymJlBcR3yxTlmzRxecG-MlasXd08WHS789829mSjGWxXvL92nWw23v37oQfotK9n0Lv8iYu_nk7tMaNbkRkcs%3D; passport_auth_mix_state=q9r0dlmb9ox9bi5rhf231vcnb7mf48qnnqfjuyze3ti49pio; IsDouyinActive=true; WallpaperGuide=%7B%22showTime%22%3A1761034790820%2C%22closeTime%22%3A0%2C%22showCount%22%3A1%2C%22cursor1%22%3A16%2C%22cursor2%22%3A4%7D; FOLLOW_LIVE_POINT_INFO=%22MS4wLjABAAAAdvkpdIfz1CAW9LFsCWeavAZa_vrx108_2f6Z817LVOI61aYpO-esUKpoIUWnbFPw%2F1761235200000%2F0%2F1761213821889%2F0%22`

	// 构造查询字符串（从curl中提取的参数）
	queryParams := map[string]string{
		"device_platform":        "webapp",
		"aid":                    "6383",
		"channel":                "channel_pc_web",
		"search_channel":         "aweme_user_web",
		"keyword":                keyword,
		"search_source":          "switch_tab",
		"query_correct_type":     "1",
		"is_filter_search":       "0",
		"from_group_id":          "",
		"disable_rs":             "0",
		"offset":                 "0",
		"count":                  "10",
		"need_filter_settings":   "1",
		"list_type":              "multi",
		"pc_search_top_1_params": "%7B%22enable_ai_search_top_1%22%3A1%7D",
		"update_version_code":    "170400",
		"pc_client_type":         "1",
		"pc_libra_divert":        "Windows",
		"support_h265":           "0",
		"support_dash":           "1",
		"cpu_core_num":           "16",
		"version_code":           "170400",
		"version_name":           "17.4.0",
		"cookie_enabled":         "true",
		"screen_width":           "2560",
		"screen_height":          "1080",
		"browser_language":       "zh-CN",
		"browser_platform":       "Win32",
		"browser_name":           "Edge",
		"browser_version":        "141.0.0.0",
		"browser_online":         "true",
		"engine_name":            "Blink",
		"engine_version":         "141.0.0.0",
		"os_name":                "Windows",
		"os_version":             "10",
		"device_memory":          "8",
		"platform":               "PC",
		"downlink":               "10",
		"effective_type":         "4g",
		"round_trip_time":        "50",
		"webid":                  "7513848749303924275",
		"uifid":                  "164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7",
		"msToken":                "rP08e-pt_nlKZOti0oLlr9cGbRskaGDO_lKy2J3lVpqO3IxB9l2d7qMBirX4PJCrdH45fWd0-uxiSjZ6sHdlXk5q_U-5z_0ZPmznKzHWxonBZXbBc6qgdC7eAJQrVLCy6ioJDL2JZx-QXM_Ph1eiQ7hmOvOUeCKy4jJZjU3X1JN7",
		//"a_bogus":                "Ey0fDzyJDNWRFdFSYKp4t1lU7lnMrBuyDBidRKCTtxulOwtOybNuMaeqaozcQY8bCYpiiC27Bj-AGnVczTUhZ9npqmZvSL0jSz2cVLvo2qpsG-khvN6zeRGNzivS0C4Yu5AVi%2FW51GMNZd5W9H9sABVHF%2F3EBRDZMr-vV%2Fujx9K4UCujw9%2F5a-bpLhVL",
	}

	// 构建查询字符串
	var paramPairs []string
	for key, value := range queryParams {
		paramPairs = append(paramPairs, fmt.Sprintf("%s=%s", key, value))
	}
	queryString := fmt.Sprintf("%s", paramPairs[0])
	for i := 1; i < len(paramPairs); i++ {
		queryString += "&" + paramPairs[i]
	}

	// 构建完整URL
	baseURL := "https://www.douyin.com/aweme/v1/web/discover/search/"
	fullURL := fmt.Sprintf("%s?%s", baseURL, queryString)

	logger.InfoWithContext(ctx, fmt.Sprintf("DouyinUserSearch 请求URL: %s", fullURL))

	// 创建请求
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return &dy, fmt.Errorf("DouyinUserSearch 创建请求失败: %v", err)
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
		return &dy, fmt.Errorf("DouyinUserSearch 请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应，处理可能的gzip压缩
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return &dy, fmt.Errorf("创建gzip读取器失败: %v", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return &dy, fmt.Errorf("DouyinUserSearch 读取响应失败: %v", err)
	}

	logger.InfoWithContext(ctx, fmt.Sprintf("DouyinUserSearch 响应数据: %s", string(body)))

	err = json.Unmarshal(body, &dy)
	if err != nil {
		return &dy, fmt.Errorf("DouyinUserSearch 解析响应JSON失败: %v", err)
	}

	return &dy, nil
}
