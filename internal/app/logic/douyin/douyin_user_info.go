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

// DouyinUserInfo 抖音用户资料信息接口
// 返回值：响应文本、HTTP 状态文本、错误
func (c *DouyinClient) DouyinUserInfo(ctx context.Context, secUserId string) (*types_douyin.DouyinUserInfoResp, error) {
	client := &http.Client{}

	logger.InfoWithContext(ctx, fmt.Sprintf("DouyinUserInfo 请求抖音用户信息: secUserId %s", secUserId))

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
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0",
		"uifid":              "164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7",
	}

	// 将 curl 的 cookie 作为整串放入 Header
	//cookieStr := `UIFID_TEMP=164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45441db5afa96ff09f5b0170e1b31058eb70a505d3baced3e44b6de3c02f2936f84; WebUgChannelId=%2230001%22; bd_ticket_guard_client_web_domain=2; SearchMultiColumnLandingAbVer=1; SEARCH_RESULT_LIST_TYPE=%22multi%22; n_mh=Hn11pbjmqhtFQ6NbHB6BtvAxooZl10kP99YFV9cRZ8g; SelfTabRedDotControl=%5B%5D; my_rd=2; live_use_vvc=%22false%22; enter_pc_once=1; d_ticket=9151353157e242473258235b66749b47d9a72; passport_assist_user=CkERozBE1KjJDCm6x2WlabcTiwnbR_tcgfuzU6kQQYe0LdCgA8GUFQQP8cvQkgh1cYK2OTENErT6aso0wiV5TH_Z0BpKCjwAAAAAAAAAAAAAT1HMv2v-0GEWxxM2txtqIiClBsZvR3gtWGiNkYTJcF4qMsFK3Kg20fVfR9_rVAm0s9QQ5dL4DRiJr9ZUIAEiAQNDpFIw; uid_tt=ef01891409b3e784d797278b4b780478; uid_tt_ss=ef01891409b3e784d797278b4b780478; sid_tt=de65f702d51d4b162ada81667eb342ad; sessionid=de65f702d51d4b162ada81667eb342ad; sessionid_ss=de65f702d51d4b162ada81667eb342ad; is_staff_user=false; login_time=1754389635201; __security_mc_1_s_sdk_cert_key=7fd9ce93-4989-bac0; __live_version__=%221.1.4.472%22; __security_mc_1_s_sdk_crypt_sdk=59402a63-4c4c-bd7f; passport_csrf_token=7f744c92495d5355cd4db2fd794e03bd; passport_csrf_token_default=7f744c92495d5355cd4db2fd794e03bd; __security_mc_1_s_sdk_sign_data_key_web_protect=65abbc52-4825-b33a; is_dash_user=1; volume_info=%7B%22isUserMute%22%3Afalse%2C%22isMute%22%3Atrue%2C%22volume%22%3A0.992%7D; sid_guard=de65f702d51d4b162ada81667eb342ad%7C1760593524%7C5184000%7CMon%2C+15-Dec-2025+05%3A45%3A24+GMT; sid_ucp_v1=1.0.0-KGIzZjc2MmE2ZjUzNzFhN2Y5ODdjODY3NGY2ZjIxZDJlNDFkNjg2MWQKIQjDoaHu5_TOAxD0jMLHBhjvMSAMMIq9n-sFOAJA8QdIBBoCaGwiIGRlNjVmNzAyZDUxZDRiMTYyYWRhODE2NjdlYjM0MmFk; ssid_ucp_v1=1.0.0-KGIzZjc2MmE2ZjUzNzFhN2Y5ODdjODY3NGY2ZjIxZDJlNDFkNjg2MWQKIQjDoaHu5_TOAxD0jMLHBhjvMSAMMIq9n-sFOAJA8QdIBBoCaGwiIGRlNjVmNzAyZDUxZDRiMTYyYWRhODE2NjdlYjM0MmFk; session_tlb_tag=sttt%7C5%7C3mX3AtUdSxYq2oFmfrNCrf________-ncq7HkaE6S1AEQwfWkJK7r67v0JN2ODxFwuBqPm8UFmA%3D; download_guide=%223%2F20251021%2F0%22; stream_player_status_params=%22%7B%5C%22is_auto_play%5C%22%3A0%2C%5C%22is_full_screen%5C%22%3A0%2C%5C%22is_full_webscreen%5C%22%3A0%2C%5C%22is_mute%5C%22%3A1%2C%5C%22is_speed%5C%22%3A1%2C%5C%22is_visible%5C%22%3A0%7D%22; publish_badge_show_info=%220%2C0%2C0%2C1761210135476%22; WallpaperGuide=%7B%22showTime%22%3A1761034790820%2C%22closeTime%22%3A0%2C%22showCount%22%3A1%2C%22cursor1%22%3A16%2C%22cursor2%22%3A4%7D; playRecommendGuideTagCount=12; totalRecommendGuideTagCount=12; strategyABtestKey=%221761269637.434%22; ttwid=1%7CCJxcNR-0EqvaaIyyL9mz09G6TkkSY1nersicZtmJqTY%7C1761269639%7C6ae9d251b197421a3d5ebd314bd51b32305a0ff8b432703f60521f80bcc1b596; FOLLOW_NUMBER_YELLOW_POINT_INFO=%22MS4wLjABAAAAdvkpdIfz1CAW9LFsCWeavAZa_vrx108_2f6Z817LVOI61aYpO-esUKpoIUWnbFPw%2F1761321600000%2F1761267189711%2F1761290598699%2F0%22; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtcmVlLXB1YmxpYy1rZXkiOiJCSk5BZzBvTkJKVFdrd3h6c1ZwVUdCaFJhUEc5aUJXL2t0RzI4aElFVGJwT2s5UjdsSkt5d01VRWlhc1E3cVpLdnpBWS81TjZXQ1J2S3JkUXY4aHJPVlE9IiwiYmQtdGlja2V0LWd1YXJkLXdlYi12ZXJzaW9uIjoyfQ%3D%3D; biz_trace_id=8ec322db; sdk_source_info=7e276470716a68645a606960273f276364697660272927676c715a6d6069756077273f276364697660272927666d776a68605a607d71606b766c6a6b5a7666776c7571273f275e58272927666a6b766a69605a696c6061273f27636469766027292762696a6764695a7364776c6467696076273f275e582729277672715a646971273f2763646976602729277f6b5a666475273f2763646976602729276d6a6e5a6b6a716c273f2763646976602729276c6b6f5a7f6367273f27636469766027292771273f2731353c3336343c3c3734333234272927676c715a75776a716a666a69273f2763646976602778; bit_env=5Gz5jSUo1c4LZxwkXadsHQpfLDZBu4_b0wVX7dYFkvxMN4E9bD2IR9V9C4DkOtzNUEw2YDVyjPQTjcHaZw24fhGrMqG5KbLy8YxphgwvG0bLYTquLl_VzSyViTmxesqdI2DS6DvaTXay3G4mobRfTxKvaUZe3rwtnmpOfyH_iXbtjF39SFM8G3ctTVXcVTdSW2cw-02m1Qj9Lb6dqAAXRbWjxSUTrUxuw2mJElT8NWxi0maIk-O8dzzL3aPUfIOBhgyKsxBdMDZGU2tyHpqx1Rk4lx2a8_wM-XCU_v0hlmZTbo4dRFHBZEPKRGY2D5qSZ8iaILr0yEsr6MXycklruM9Y0INDY-MzQwu4VT6qPffwBFrlZKq_sY3AIKUo7NndgBriW9Xck0HfduC0MpBnwFBO2GIPvzjmurpyMoZxmXaMN30LQ0Cf14F1fcxjJNep4M2KVdK44BuvThzWtGVK-o1tarqXNbztwpTrnaBMaVNT_Be56iBx0iZMazvX6urNlHeI0PfqLi3spH03a0sTiQHYI-AB5eHDSLZzNQj0Hhs%3D; gulu_source_res=eyJwX2luIjoiZjI1NzFkMzg0MDZkYWFhM2I1MGFkY2E0MjgxMDI4N2VmMDEwMDcxYjQzNTA2ZWJkY2RlOGYxZDZmMjYyZWQ0NCJ9; passport_auth_mix_state=hvbqvfwg4e5qvg4369ewck0mkuxprjzhi3llb5scev3hibi4; odin_tt=6ae94e823fceae68676fbaad0c6f3fc348f0b24011ccbd64c840cfb7fbb108fa90cab9e32356ea85f7eff2b0b8cc2e16881bae8eb6edba2316e039d08c3f9244; bd_ticket_guard_client_data_v2=eyJyZWVfcHVibGljX2tleSI6IkJKTkFnMG9OQkpUV2t3eHpzVnBVR0JoUmFQRzlpQlcva3RHMjhoSUVUYnBPazlSN2xKS3l3TVVFaWFzUTdxWkt2ekFZLzVONldDUnZLcmRRdjhock9WUT0iLCJ0c19zaWduIjoidHMuMi4zNjE4YzJlZmY1OGNmOWMwZmJmNzg3ZDAwZDUwMjYxZWVmZGQ3MDc4MjE1ZDg1ZmYxYWZhNjc4MWI3OTMwMTc4YzRmYmU4N2QyMzE5Y2YwNTMxODYyNGNlZGExNDkxMWNhNDA2ZGVkYmViZWRkYjJlMzBmY2U4ZDRmYTAyNTc1ZCIsInJlcV9jb250ZW50Ijoic2VjX3RzIiwicmVxX3NpZ24iOiJ3bE5Cdi84cWR0RHVWUHJjTU40Ny8vWWFhSUt1WU5tNmRwSEE1TTJmdnhNPSIsInNlY190cyI6IiNWYnFoTzZzNUpDVHE1SDBZZlE3NEZLK2pzVTBQQlk5QXcrSzRERzVIMlpES3gybFRxQUtKeldUQ2hGVTYifQ%3D%3D; IsDouyinActive=true; home_can_add_dy_2_desktop=%220%22; stream_recommend_feed_params=%22%7B%5C%22cookie_enabled%5C%22%3Atrue%2C%5C%22screen_width%5C%22%3A2560%2C%5C%22screen_height%5C%22%3A1080%2C%5C%22browser_online%5C%22%3Atrue%2C%5C%22cpu_core_num%5C%22%3A16%2C%5C%22device_memory%5C%22%3A8%2C%5C%22downlink%5C%22%3A10%2C%5C%22effective_type%5C%22%3A%5C%224g%5C%22%2C%5C%22round_trip_time%5C%22%3A50%7D%22; FOLLOW_LIVE_POINT_INFO=%22MS4wLjABAAAAdvkpdIfz1CAW9LFsCWeavAZa_vrx108_2f6Z817LVOI61aYpO-esUKpoIUWnbFPw%2F1761321600000%2F0%2F1761299336186%2F0%22`

	// 构造查询字符串（顺序与 curl 保持一致，不包含 a_bogus）
	pa := fmt.Sprintf("device_platform=%s&aid=%s&channel=%s&publish_video_strategy_type=%s&source=%s&sec_user_id=%s&personal_center_strategy=%s&profile_other_record_enable=%s&land_to=%s&update_version_code=%s&pc_client_type=%s&pc_libra_divert=%s&support_h265=%s&support_dash=%s&cpu_core_num=%s&version_code=%s&version_name=%s&cookie_enabled=%s&screen_width=%s&screen_height=%s&browser_language=%s&browser_platform=%s&browser_name=%s&browser_version=%s&browser_online=%s&engine_name=%s&engine_version=%s&os_name=%s&os_version=%s&device_memory=%s&platform=%s&downlink=%s&effective_type=%s&round_trip_time=%s&webid=%s&uifid=%s&msToken=%s&verifyFp=%s&fp=%s",
		url.QueryEscape("webapp"),
		url.QueryEscape("6383"),
		url.QueryEscape("channel_pc_web"),
		url.QueryEscape("2"),
		url.QueryEscape("channel_pc_web"),
		url.QueryEscape(secUserId),
		url.QueryEscape("1"),
		url.QueryEscape("1"),
		url.QueryEscape("1"),
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
		url.QueryEscape("50"),
		url.QueryEscape("7513848749303924275"),
		url.QueryEscape("164c22db5016193fd69c8bfb0b166ea3a563c2c88054b8eae8759946ea9753ce12cc8cddeedd8cda7f6e9d87be58e45488d21356233118598c8bd522378501089b006c79d3d16665745052337acd9d2a908bb62bf139a4a48adb3cabe85ac607c2faf710ec3457be6c2fd3b0a5047945a7b237c5971934276ad04f274fca7731d8bf3998d89071b9a8dc98bf8971a18f7eca4f76ceace138716a4cd155eecec7"),
		url.QueryEscape("Ry7hniBwri-dEsnMD4TLu07MWx--9ij1cgwSbSPHH7_i34GzTvlV-NVC2vIlU3yJUczIS4eU8k2Wu4nBl6OwkHAKJrRY6fq0KiBQVhRZH7WV6MXOI4GNacmtVlQl5dj0V_GSqK4YGskYD-YSPKC6MEXRZN3WGE-nNeFuOBkkFZQcuQ5vlUDD2Q%3D%3D"),
		url.QueryEscape("verify_mgissah8_C6nEvcyk_5pgr_4Qb6_9BGQ_phmZxr733zQ4"),
		url.QueryEscape("verify_mgissah8_C6nEvcyk_5pgr_4Qb6_9BGQ_phmZxr733zQ4"),
	)

	// GET 无请求体，da 为空字符串
	//da := ""

	// 生成 a_bogus
	//aBogus, err := c.generateABogusV2(pa, da)
	//if err != nil {
	//	return nil, "", fmt.Errorf("生成a_bogus失败: %v", err)
	//}

	// 构建完整URL
	baseURL := "https://www-hj.douyin.com/aweme/v1/web/user/profile/other/"
	fullURL := fmt.Sprintf("%s?%s", baseURL, pa)

	logger.InfoWithContext(ctx, fmt.Sprintf("DouyinUserInfo 请求URL: %s", fullURL))
	// 创建请求
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("DouyinUserInfo 创建请求失败: %v", err)
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
		return nil, fmt.Errorf("DouyinUserInfo 请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应，处理可能的gzip压缩
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("创建gzip读取器失败: %v", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	logger.InfoWithContext(ctx, fmt.Sprintf("DouyinUserInfo 响应: %s", string(body)))
	var dyInfo types_douyin.DouyinUserInfoResp
	json.Unmarshal(body, &dyInfo)

	return &dyInfo, nil
}
