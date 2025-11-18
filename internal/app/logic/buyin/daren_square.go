package buyin

import (
	"encoding/json"
	"fmt"
)

// FiltersData 达人广场筛选条件数据结构
type FiltersData struct {
	MainCateNew                         []string `json:"main_cate_new"`
	ContentType                         []string `json:"content_type"`
	TotalSales                          []string `json:"total_sales"`
	LiveTotalSales                      []string `json:"live_total_sales"`
	VideoTotalSales                     []string `json:"video_total_sales"`
	ImageTextTotalSales                 []string `json:"image_text_total_sales"`
	WindowsTotalSales                   []string `json:"windows_total_sales"`
	LiveSaleAvg                         []string `json:"live_sale_avg"`
	LiveWatchingTimes                   []string `json:"live_watching_times"`
	VideoSaleSingle                     []string `json:"video_sale_single"`
	VideoPlayTime                       []string `json:"video_play_time"`
	CommonRangeSelectionVideoNatureRate []string `json:"common_range_selection_video_nature_rate"`
	ImageTextSaleSingle                 []string `json:"image_text_sale_single"`
	ImageTextPlayTime                   []string `json:"image_text_play_time"`
	WindowOrderCnt                      []string `json:"window_order_cnt"`
	WindowProductCnt                    []string `json:"window_product_cnt"`
	AuthorLevelNew                      []string `json:"author_level_new"`
	FansNum                             []string `json:"fans_num"`
	AuthorPortrait                      []string `json:"author_portrait"`
	FanPortrait                         []string `json:"fan_portrait"`
	FansProfile                         []string `json:"fans_profile"`
	HighResponseRateIm                  []string `json:"high_response_rate_im"`
	HasContact                          []string `json:"has_contact"`
	CommonSelectionNoViewContact60d     []string `json:"common_selection_no_view_contact_60d"`
	CanConnect                          []string `json:"can_connect"`
	CommonSelectionNoSendImMessage60d   []string `json:"common_selection_no_send_im_message_60d"`
	CommonSelectionAuthorAccountType    []string `json:"common_selection_author_account_type"`
	CommonSelectionHideNewsAccount      []string `json:"common_selection_hide_news_account"`
}

// DarenSquareRequest 达人广场请求参数
type DarenSquareRequest struct {
	Page     int         `json:"page"`
	Refresh  bool        `json:"refresh"`
	Type     int         `json:"type"`
	SearchID string      `json:"search_id"`
	Query    string      `json:"query"`
	Filters  FiltersData `json:"filters"`
}

// SearchFeedAuthor 搜索达人广场
func (c *JXClient) SearchFeedAuthor(requestData DarenSquareRequest) (*FeedAuthorList, error) {
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"content-type":       "application/json",
		"origin":             "https://buyin.jinritemai.com",
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

	queryParams := map[string]string{
		"ewid":     c.config.EWID,
		"verifyFp": c.config.VerifyFp,
		"fp":       c.config.Fp,
		"msToken":  c.config.MsToken,
	}

	// 序列化请求体
	body, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求数据失败: %v", err)
	}

	baseURL := "https://buyin.jinritemai.com/square_pc_api/square/search_feed_author"
	request, _, err := c.doRequest("POST", baseURL, headers, queryParams, body)

	var resp FeedAuthorList
	json.Unmarshal([]byte(request), &resp)
	if resp.Code != 0 {
		return nil, fmt.Errorf("响应错误: %s", resp.Msg)
	}
	return &resp, err
}

type FeedAuthorList struct {
	Code int `json:"code"`
	Data struct {
		HasMore   bool `json:"has_more"`
		ResSource int  `json:"res_source"`
		List      []struct {
			AuthorBase struct {
				Uid         string `json:"uid"`
				AuthorId    int    `json:"author_id"`
				Nickname    string `json:"nickname"`
				Avatar      string `json:"avatar"`
				FansNum     int    `json:"fans_num"`
				Gender      int    `json:"gender"`
				City        string `json:"city"`
				AuthorLevel int    `json:"author_level"`
				AvatarBig   string `json:"avatar_big"`
				AwemeId     string `json:"aweme_id"`
			} `json:"author_base"`
			AuthorTag struct {
				MainCate           []string `json:"main_cate"`
				DarkHorse          string   `json:"dark_horse"`
				ContactIcon        string   `json:"contact_icon"`
				HighReply          string   `json:"high_reply"`
				InvitationStatus   int      `json:"invitation_status"`
				InviteStatus       int      `json:"invite_status"`
				SatisfyRequirement int      `json:"satisfy_requirement"`
				AlreadyCooperated  bool     `json:"already_cooperated"`
				IsStar             bool     `json:"is_star"`
				HighResponseRate   string   `json:"high_response_rate"`
				IsOpenInvoice      bool     `json:"is_open_invoice"`
				AuthorRecReasons   []struct {
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
				AuthorLabelRecReasons []struct {
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
				} `json:"author_label_rec_reasons"`
			} `json:"author_tag"`
			AuthorContact struct {
				Phone  string `json:"phone"`
				Wechat string `json:"wechat"`
				Lark   string `json:"lark"`
				Douyin string `json:"douyin"`
			} `json:"author_contact"`
			AuthorLive struct {
				WatchingNumber int `json:"watching_number"`
				SaleStatus     int `json:"sale_status"`
				SaleLow        int `json:"sale_low"`
				SaleHigh       int `json:"sale_high"`
				GPMStatus      int `json:"GPM_status"`
				GPMLow         int `json:"GPM_low"`
				GPMHigh        int `json:"GPM_high"`
				WatchingTimes  int `json:"watching_times"`
				AllLiveNum30D  int `json:"all_live_num_30d"`
			} `json:"author_live"`
			AuthorVideo struct {
				PlayMedian      int `json:"play_median"`
				GPMStatus       int `json:"GPM_status"`
				GPMLow          int `json:"GPM_low"`
				GPMHigh         int `json:"GPM_high"`
				VideoSaleStatus int `json:"video_sale_status"`
				VideoSaleLow    int `json:"video_sale_low"`
				VideoSaleHigh   int `json:"video_sale_high"`
				WatchingTimes   int `json:"watching_times"`
				AllVideoNum30D  int `json:"all_video_num_30d"`
			} `json:"author_video"`
			AuthorConf struct {
				SaleRequirement struct {
					Type      int    `json:"type"`
					Value     string `json:"value"`
					PriceType int    `json:"price_type"`
				} `json:"sale_requirement"`
			} `json:"author_conf"`
			AuthorScore struct {
				ExprScore struct {
					Score      int    `json:"score"`
					Level      string `json:"level"`
					Percentage int    `json:"percentage"`
				} `json:"expr_score"`
				ProductScore struct {
					Score      int    `json:"score"`
					Level      string `json:"level"`
					Percentage int    `json:"percentage"`
				} `json:"product_score"`
				LogisticsScore struct {
					Score      int    `json:"score"`
					Level      string `json:"level"`
					Percentage int    `json:"percentage"`
				} `json:"logistics_score"`
				ServiceScore struct {
					Score      int    `json:"score"`
					Level      string `json:"level"`
					Percentage int    `json:"percentage"`
				} `json:"service_score"`
			} `json:"author_score"`
			AuthorSale struct {
				SaleStatus          int    `json:"sale_status"`
				SaleD30Low          int    `json:"sale_d30_low"`
				SaleD30High         int    `json:"sale_d30_high"`
				MainSaleType        string `json:"main_sale_type"`
				PromotionNum        int    `json:"promotion_num"`
				CooperateContentNum int    `json:"cooperate_content_num"`
				CooperateProductNum int    `json:"cooperate_product_num"`
			} `json:"author_sale"`
			SaleInfo struct {
				TotalSales struct {
					SaleStatus int `json:"sale_status"`
					SaleValue  int `json:"sale_value"`
					SaleLow    int `json:"sale_low"`
					SaleHigh   int `json:"sale_high"`
				} `json:"total_sales"`
				LiveTotalSales struct {
					SaleStatus int `json:"sale_status"`
					SaleValue  int `json:"sale_value"`
					SaleLow    int `json:"sale_low"`
					SaleHigh   int `json:"sale_high"`
				} `json:"live_total_sales"`
				VideoTotalSales struct {
					SaleStatus int `json:"sale_status"`
					SaleValue  int `json:"sale_value"`
					SaleLow    int `json:"sale_low"`
					SaleHigh   int `json:"sale_high"`
				} `json:"video_total_sales"`
				ImageTextTotalSales struct {
					SaleStatus int `json:"sale_status"`
					SaleValue  int `json:"sale_value"`
					SaleLow    int `json:"sale_low"`
					SaleHigh   int `json:"sale_high"`
				} `json:"image_text_total_sales"`
				WindowTotalSales struct {
					SaleStatus int `json:"sale_status"`
					SaleValue  int `json:"sale_value"`
					SaleLow    int `json:"sale_low"`
					SaleHigh   int `json:"sale_high"`
				} `json:"window_total_sales"`
			} `json:"sale_info"`
			AuthorImageText struct {
				ImageTextNum       int `json:"image_text_num"`
				WatchingTimes      int `json:"watching_times"`
				GPMStatus          int `json:"GPM_status"`
				GPMLow             int `json:"GPM_low"`
				GPMHigh            int `json:"GPM_high"`
				SaleLow            int `json:"sale_low"`
				SaleHigh           int `json:"sale_high"`
				AllImageTextNum30D int `json:"all_image_text_num_30d"`
				SaleStatus         int `json:"sale_status"`
			} `json:"author_image_text"`
			AuthorWindow struct {
				ProductNum int `json:"product_num"`
				OrderLow   int `json:"order_low"`
				OrderHigh  int `json:"order_high"`
				Status     int `json:"status"`
			} `json:"author_window"`
			AuthorCrm struct {
				IsJoin bool `json:"is_join"`
			} `json:"author_crm"`
			AuthorConnection struct {
				RuleLink    string `json:"rule_link"`
				ImStatus    int    `json:"im_status"`
				ImHoverText string `json:"im_hover_text"`
			} `json:"author_connection"`
		} `json:"list"`
		Total           int    `json:"total"`
		EventTrackLogId string `json:"event_track_log_id"`
	} `json:"data"`
	LogId string `json:"log_id"`
	Msg   string `json:"msg"`
	St    int    `json:"st"`
}
