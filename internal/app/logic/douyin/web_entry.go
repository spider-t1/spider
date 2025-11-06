package douyin

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DouyinLiveRoomEnter 调用直播间进入接口（webcast/room/web/enter）
// 参考 web_entry.txt 的 curl 参数与头信息
func (c *DouyinClient) DouyinLiveRoomEnter(ctx context.Context, webRid string) (*LiveRoomEnterResp, error) {
	client := &http.Client{}

	//logger.InfoWithContext(ctx, fmt.Sprintf("DouyinLiveRoomEnter 请求参数: webRid=%s roomIdStr=%s", , ))

	// 请求头（参考 curl）
	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"priority":           "u=1, i",
		"referer":            fmt.Sprintf("https://live.douyin.com/%s", webRid),
		"sec-ch-ua":          "\"Chromium\";v=\"142\", \"Microsoft Edge\";v=\"142\", \"Not_A Brand\";v=\"99\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         c.UserAgent,
	}

	// 构造查询字符串（顺序与 curl 保持一致，不包含 a_bogus）
	pa := fmt.Sprintf("aid=%s&app_name=%s&live_id=%s&device_platform=%s&language=%s&enter_from=%s&cookie_enabled=%s&screen_width=%s&screen_height=%s&browser_language=%s&browser_platform=%s&browser_name=%s&browser_version=%s&web_rid=%s&room_id_str=%s&enter_source=%s&is_need_double_stream=%s&insert_task_id=%s&live_reason=%s&msToken=%s",
		url.QueryEscape("6383"),
		url.QueryEscape("douyin_web"),
		url.QueryEscape("1"),
		url.QueryEscape("web"),
		url.QueryEscape("zh-CN"),
		url.QueryEscape("web_search"),
		url.QueryEscape("true"),
		url.QueryEscape("2560"),
		url.QueryEscape("1080"),
		url.QueryEscape("zh-CN"),
		url.QueryEscape("Win32"),
		url.QueryEscape("Edge"),
		url.QueryEscape("142.0.0.0"),
		url.QueryEscape(webRid),
		url.QueryEscape(""),
		url.QueryEscape(""),
		url.QueryEscape("false"),
		url.QueryEscape(""),
		url.QueryEscape(""),
		url.QueryEscape("-XYoRSIj007_Cne-0M-0mIck5LyOGu1U7EH_26I6AaCjN5G-_ETohUeYyZiUBooGdJe6EkK70rZHWnvSP6MTnrBr-XwmCkokSmddoQmGUbk4nojnEZI-uR4M5tMrlzqLVqN48uvSuwzZlBmAXFgNmW82xoG6MPeZmVHOyeL5HLdSPnQ1P23X3g%3D%3D"),
	)

	// GET 无请求体，da 为空字符串
	da := ""

	// 生成 a_bogus
	aBogus, err := c.generateABogusV2(pa, da)
	if err != nil {
		return nil, fmt.Errorf("生成a_bogus失败: %v", err)
	}

	// 构建完整URL
	baseURL := "https://live.douyin.com/webcast/room/web/enter/"
	fullURL := fmt.Sprintf("%s?%s&a_bogus=%s", baseURL, pa, url.QueryEscape(aBogus))

	//logger.InfoWithContext(ctx, fmt.Sprintf("DouyinLiveRoomEnter 请求URL: %s", fullURL))

	// 创建请求
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
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
		return nil, fmt.Errorf("请求失败: %v", err)
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

	//logger.InfoWithContext(ctx, fmt.Sprintf("DouyinLiveRoomEnter 响应: %s", string(body)))

	var out LiveRoomEnterResp
	if err = json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("解析响应JSON失败: %v", err)
	}

	return &out, nil
}

type LiveRoomEnterResp struct {
	Data struct {
		Data []struct {
			IdStr        string `json:"id_str"`
			Status       int    `json:"status"`
			StatusStr    string `json:"status_str"`
			Title        string `json:"title"`
			UserCountStr string `json:"user_count_str"`
			Cover        struct {
				UrlList []string `json:"url_list"`
			} `json:"cover"`
			StreamUrl struct {
				FlvPullUrl struct {
					FULLHD1 string `json:"FULL_HD1"`
					HD1     string `json:"HD1"`
					SD1     string `json:"SD1"`
					SD2     string `json:"SD2"`
				} `json:"flv_pull_url"`
				DefaultResolution string `json:"default_resolution"`
				HlsPullUrlMap     struct {
					FULLHD1 string `json:"FULL_HD1"`
					HD1     string `json:"HD1"`
					SD1     string `json:"SD1"`
					SD2     string `json:"SD2"`
				} `json:"hls_pull_url_map"`
				HlsPullUrl        string `json:"hls_pull_url"`
				StreamOrientation int    `json:"stream_orientation"`
				LiveCoreSdkData   struct {
					PullData struct {
						Options struct {
							DefaultQuality struct {
								Name              string `json:"name"`
								SdkKey            string `json:"sdk_key"`
								VCodec            string `json:"v_codec"`
								Resolution        string `json:"resolution"`
								Level             int    `json:"level"`
								VBitRate          int    `json:"v_bit_rate"`
								AdditionalContent string `json:"additional_content"`
								Fps               int    `json:"fps"`
								Disable           int    `json:"disable"`
							} `json:"default_quality"`
							Qualities []struct {
								Name              string `json:"name"`
								SdkKey            string `json:"sdk_key"`
								VCodec            string `json:"v_codec"`
								Resolution        string `json:"resolution"`
								Level             int    `json:"level"`
								VBitRate          int    `json:"v_bit_rate"`
								AdditionalContent string `json:"additional_content"`
								Fps               int    `json:"fps"`
								Disable           int    `json:"disable"`
							} `json:"qualities"`
						} `json:"options"`
						StreamData string `json:"stream_data"`
					} `json:"pull_data"`
				} `json:"live_core_sdk_data"`
				Extra struct {
					Height                  int  `json:"height"`
					Width                   int  `json:"width"`
					Fps                     int  `json:"fps"`
					MaxBitrate              int  `json:"max_bitrate"`
					MinBitrate              int  `json:"min_bitrate"`
					DefaultBitrate          int  `json:"default_bitrate"`
					BitrateAdaptStrategy    int  `json:"bitrate_adapt_strategy"`
					AnchorInteractProfile   int  `json:"anchor_interact_profile"`
					AudienceInteractProfile int  `json:"audience_interact_profile"`
					HardwareEncode          bool `json:"hardware_encode"`
					VideoProfile            int  `json:"video_profile"`
					H265Enable              bool `json:"h265_enable"`
					GopSec                  int  `json:"gop_sec"`
					BframeEnable            bool `json:"bframe_enable"`
					Roi                     bool `json:"roi"`
					SwRoi                   bool `json:"sw_roi"`
					Bytevc1Enable           bool `json:"bytevc1_enable"`
				} `json:"extra"`
				PullDatas struct {
				} `json:"pull_datas"`
			} `json:"stream_url"`
			MosaicStatus    int      `json:"mosaic_status"`
			MosaicStatusStr string   `json:"mosaic_status_str"`
			AdminUserIds    []int64  `json:"admin_user_ids"`
			AdminUserIdsStr []string `json:"admin_user_ids_str"`
			Owner           struct {
				IdStr       string `json:"id_str"`
				SecUid      string `json:"sec_uid"`
				Nickname    string `json:"nickname"`
				AvatarThumb struct {
					UrlList []string `json:"url_list"`
				} `json:"avatar_thumb"`
				FollowInfo struct {
					FollowStatus    int    `json:"follow_status"`
					FollowStatusStr string `json:"follow_status_str"`
				} `json:"follow_info"`
				Subscribe struct {
					IsMember     bool `json:"is_member"`
					Level        int  `json:"level"`
					IdentityType int  `json:"identity_type"`
					BuyType      int  `json:"buy_type"`
					Open         int  `json:"open"`
				} `json:"subscribe"`
				ForeignUser int    `json:"foreign_user"`
				OpenIdStr   string `json:"open_id_str"`
			} `json:"owner"`
			RoomAuth struct {
				Chat                      bool `json:"Chat"`
				Danmaku                   bool `json:"Danmaku"`
				Gift                      bool `json:"Gift"`
				LuckMoney                 bool `json:"LuckMoney"`
				Digg                      bool `json:"Digg"`
				RoomContributor           bool `json:"RoomContributor"`
				Props                     bool `json:"Props"`
				UserCard                  bool `json:"UserCard"`
				POI                       bool `json:"POI"`
				MoreAnchor                int  `json:"MoreAnchor"`
				Banner                    int  `json:"Banner"`
				Share                     int  `json:"Share"`
				UserCorner                int  `json:"UserCorner"`
				Landscape                 int  `json:"Landscape"`
				LandscapeChat             int  `json:"LandscapeChat"`
				PublicScreen              int  `json:"PublicScreen"`
				GiftAnchorMt              int  `json:"GiftAnchorMt"`
				RecordScreen              int  `json:"RecordScreen"`
				DonationSticker           int  `json:"DonationSticker"`
				HourRank                  int  `json:"HourRank"`
				CommerceCard              int  `json:"CommerceCard"`
				AudioChat                 int  `json:"AudioChat"`
				DanmakuDefault            int  `json:"DanmakuDefault"`
				KtvOrderSong              int  `json:"KtvOrderSong"`
				SelectionAlbum            int  `json:"SelectionAlbum"`
				Like                      int  `json:"Like"`
				MultiplierPlayback        int  `json:"MultiplierPlayback"`
				DownloadVideo             int  `json:"DownloadVideo"`
				Collect                   int  `json:"Collect"`
				TimedShutdown             int  `json:"TimedShutdown"`
				Seek                      int  `json:"Seek"`
				Denounce                  int  `json:"Denounce"`
				Dislike                   int  `json:"Dislike"`
				OnlyTa                    int  `json:"OnlyTa"`
				CastScreen                int  `json:"CastScreen"`
				CommentWall               int  `json:"CommentWall"`
				BulletStyle               int  `json:"BulletStyle"`
				ShowGamePlugin            int  `json:"ShowGamePlugin"`
				VSGift                    int  `json:"VSGift"`
				VSTopic                   int  `json:"VSTopic"`
				VSRank                    int  `json:"VSRank"`
				AdminCommentWall          int  `json:"AdminCommentWall"`
				CommerceComponent         int  `json:"CommerceComponent"`
				DouPlus                   int  `json:"DouPlus"`
				GamePointsPlaying         int  `json:"GamePointsPlaying"`
				Poster                    int  `json:"Poster"`
				Highlights                int  `json:"Highlights"`
				TypingCommentState        int  `json:"TypingCommentState"`
				StrokeUpDownGuide         int  `json:"StrokeUpDownGuide"`
				UpRightStatsFloatingLayer int  `json:"UpRightStatsFloatingLayer"`
				CastScreenExplicit        int  `json:"CastScreenExplicit"`
				Selection                 int  `json:"Selection"`
				IndustryService           int  `json:"IndustryService"`
				VerticalRank              int  `json:"VerticalRank"`
				EnterEffects              int  `json:"EnterEffects"`
				FansClub                  int  `json:"FansClub"`
				EmojiOutside              int  `json:"EmojiOutside"`
				CanSellTicket             int  `json:"CanSellTicket"`
				DouPlusPopularityGem      int  `json:"DouPlusPopularityGem"`
				MissionCenter             int  `json:"MissionCenter"`
				ExpandScreen              int  `json:"ExpandScreen"`
				FansGroup                 int  `json:"FansGroup"`
				Topic                     int  `json:"Topic"`
				AnchorMission             int  `json:"AnchorMission"`
				Teleprompter              int  `json:"Teleprompter"`
				LongTouch                 int  `json:"LongTouch"`
				FirstFeedHistChat         int  `json:"FirstFeedHistChat"`
				MoreHistChat              int  `json:"MoreHistChat"`
				TaskBanner                int  `json:"TaskBanner"`
				SpecialStyle              struct {
					Chat struct {
						UnableStyle             int    `json:"UnableStyle"`
						Content                 string `json:"Content"`
						OffType                 int    `json:"OffType"`
						AnchorSwitchForPaidLive int    `json:"AnchorSwitchForPaidLive"`
						ContentForPaidLive      string `json:"ContentForPaidLive"`
					} `json:"Chat"`
					Like struct {
						UnableStyle             int    `json:"UnableStyle"`
						Content                 string `json:"Content"`
						OffType                 int    `json:"OffType"`
						AnchorSwitchForPaidLive int    `json:"AnchorSwitchForPaidLive"`
						ContentForPaidLive      string `json:"ContentForPaidLive"`
					} `json:"Like"`
				} `json:"SpecialStyle"`
				FixedChat             int `json:"FixedChat"`
				QuizGamePointsPlaying int `json:"QuizGamePointsPlaying"`
			} `json:"room_auth"`
			LiveRoomMode int `json:"live_room_mode"`
			Stats        struct {
				TotalUserDesp string `json:"total_user_desp"`
				LikeCount     int    `json:"like_count"`
				TotalUserStr  string `json:"total_user_str"`
				UserCountStr  string `json:"user_count_str"`
			} `json:"stats"`
			HasCommerceGoods bool `json:"has_commerce_goods"`
			LinkerMap        struct {
			} `json:"linker_map"`
			LinkerDetail struct {
				LinkerPlayModes             []interface{} `json:"linker_play_modes"`
				BigPartyLayoutConfigVersion int           `json:"big_party_layout_config_version"`
				AcceptAudiencePreApply      bool          `json:"accept_audience_pre_apply"`
				LinkerUiLayout              int           `json:"linker_ui_layout"`
				EnableAudienceLinkmic       int           `json:"enable_audience_linkmic"`
				FunctionType                string        `json:"function_type"`
				LinkerMapStr                struct {
				} `json:"linker_map_str"`
				KtvLyricMode             string `json:"ktv_lyric_mode"`
				InitSource               string `json:"init_source"`
				ForbidApplyFromOther     bool   `json:"forbid_apply_from_other"`
				KtvExhibitMode           int    `json:"ktv_exhibit_mode"`
				EnlargeGuestTurnOnSource int    `json:"enlarge_guest_turn_on_source"`
				PlaymodeDetail           struct {
				} `json:"playmode_detail"`
				ClientUiInfo string        `json:"client_ui_info"`
				ManualOpenUi int           `json:"manual_open_ui"`
				FeatureList  []interface{} `json:"feature_list"`
			} `json:"linker_detail"`
			RoomViewStats struct {
				IsHidden            bool   `json:"is_hidden"`
				DisplayShort        string `json:"display_short"`
				DisplayMiddle       string `json:"display_middle"`
				DisplayLong         string `json:"display_long"`
				DisplayValue        int    `json:"display_value"`
				DisplayVersion      int    `json:"display_version"`
				Incremental         bool   `json:"incremental"`
				DisplayType         int    `json:"display_type"`
				DisplayShortAnchor  string `json:"display_short_anchor"`
				DisplayMiddleAnchor string `json:"display_middle_anchor"`
				DisplayLongAnchor   string `json:"display_long_anchor"`
			} `json:"room_view_stats"`
			SceneTypeInfo struct {
				IsUnionLiveRoom              bool `json:"is_union_live_room"`
				IsLife                       bool `json:"is_life"`
				IsProtectedRoom              int  `json:"is_protected_room"`
				IsLastedGoodsRoom            int  `json:"is_lasted_goods_room"`
				IsDesireRoom                 int  `json:"is_desire_room"`
				CommentaryType               bool `json:"commentary_type"`
				IsSubOrientationVerticalRoom int  `json:"is_sub_orientation_vertical_room"`
			} `json:"scene_type_info"`
			ToolbarData struct {
				EntranceList []struct {
					GroupId       int    `json:"group_id"`
					ComponentType int    `json:"component_type"`
					OpType        int    `json:"op_type"`
					Text          string `json:"text"`
					SchemaUrl     string `json:"schema_url"`
					ShowType      int    `json:"show_type"`
					DataStatus    int    `json:"data_status"`
					Extra         string `json:"extra"`
					Icon          struct {
						UrlList         []string      `json:"url_list"`
						Uri             string        `json:"uri"`
						Height          int           `json:"height"`
						Width           int           `json:"width"`
						AvgColor        string        `json:"avg_color"`
						ImageType       int           `json:"image_type"`
						OpenWebUrl      string        `json:"open_web_url"`
						IsAnimated      bool          `json:"is_animated"`
						FlexSettingList []interface{} `json:"flex_setting_list"`
						TextSettingList []interface{} `json:"text_setting_list"`
					} `json:"icon,omitempty"`
				} `json:"entrance_list"`
				MorePanel []struct {
					GroupId       int    `json:"group_id"`
					ComponentType int    `json:"component_type"`
					OpType        int    `json:"op_type"`
					Text          string `json:"text"`
					SchemaUrl     string `json:"schema_url"`
					ShowType      int    `json:"show_type"`
					DataStatus    int    `json:"data_status"`
					Extra         string `json:"extra"`
				} `json:"more_panel"`
				MaxEntranceCnt   int           `json:"max_entrance_cnt"`
				LandscapeUpRight []interface{} `json:"landscape_up_right"`
				SkinResource     struct {
				} `json:"skin_resource"`
				MaxEntranceCntLandscape int `json:"max_entrance_cnt_landscape"`
				Permutation             struct {
					General struct {
						GroupPriority     []int `json:"GroupPriority"`
						ComponentSequence []int `json:"ComponentSequence"`
					} `json:"general"`
					OnDemandComponentList []interface{} `json:"on_demand_component_list"`
					BlockComponentList    []interface{} `json:"block_component_list"`
				} `json:"permutation"`
				ExtraInfo struct {
					GamePromotionCoexist int `json:"game_promotion_coexist"`
				} `json:"extra_info"`
			} `json:"toolbar_data"`
			EcomData struct {
				RedsShowInfos []interface{} `json:"reds_show_infos"`
				InstantType   int           `json:"instant_type"`
				RouteRule     string        `json:"route_rule"`
			} `json:"ecom_data"`
			RoomCart struct {
				ContainCart bool   `json:"contain_cart"`
				Total       int    `json:"total"`
				FlashTotal  int    `json:"flash_total"`
				CartIcon    string `json:"cart_icon"`
				ShowCart    int    `json:"show_cart"`
			} `json:"room_cart"`
			AnchorABMap struct {
				AbAdminCommentOnWall                string `json:"ab_admin_comment_on_wall"`
				AbFriendChat                        string `json:"ab_friend_chat"`
				AdminOptimizeThird                  string `json:"admin_optimize_third"`
				AdminPrivilegeRefine                string `json:"admin_privilege_refine"`
				AllowSharedToFans                   string `json:"allow_shared_to_fans"`
				AudienceLinkmicContinue             string `json:"audience_linkmic_continue"`
				Audio1V8StageEnlarge                string `json:"audio_1v8_stage_enlarge"`
				AudioDoubleEnlargeEnable            string `json:"audio_double_enlarge_enable"`
				AudioRadioV2                        string `json:"audio_radio_v2"`
				AudioRoomSubtitleOpt                string `json:"audio_room_subtitle_opt"`
				BattleMatchRebuildAnchor            string `json:"battle_match_rebuild_anchor"`
				BigPartyEnableOpenCamera            string `json:"big_party_enable_open_camera"`
				ChatIntercommunicateMultiAnchor     string `json:"chat_intercommunicate_multi_anchor"`
				ChatIntercommunicatePk              string `json:"chat_intercommunicate_pk"`
				CrossDefaultEnlarge                 string `json:"cross_default_enlarge"`
				CrossLinkSupportEnlargeGuest        string `json:"cross_link_support_enlarge_guest"`
				CrossRoomBattlePopMode              string `json:"cross_room_battle_pop_mode"`
				DoubleEnlargeEnable                 string `json:"double_enlarge_enable"`
				EcomRoomDisableGift                 string `json:"ecom_room_disable_gift"`
				EnableEnterBySharing                string `json:"enable_enter_by_sharing"`
				EnableLinkGuestEnter                string `json:"enable_link_guest_enter"`
				EnterMessageTipRelation             string `json:"enter_message_tip_relation"`
				EnterSourceMark                     string `json:"enter_source_mark"`
				FrequentlyChatAbValue               string `json:"frequently_chat_ab_value"`
				FriendRoomAudioTuning               string `json:"friend_room_audio_tuning"`
				FriendRoomSupportNsMode             string `json:"friend_room_support_ns_mode"`
				FriendShareVideoFeatureType         string `json:"friend_share_video_feature_type"`
				GameLinkEntrance                    string `json:"game_link_entrance"`
				GiftComment                         string `json:"gift_comment"`
				GiftCommentV2                       string `json:"gift_comment_v2"`
				GiftHideTip                         string `json:"gift_hide_tip"`
				GuestBattleCrownUpgrade             string `json:"guest_battle_crown_upgrade"`
				GuestBattleExpand                   string `json:"guest_battle_expand"`
				GuestBattleScoreExpand              string `json:"guest_battle_score_expand"`
				GuestBattleUpgrade                  string `json:"guest_battle_upgrade"`
				InteractActingAb                    string `json:"interact_acting_ab"`
				InteractAnchorGuide                 string `json:"interact_anchor_guide"`
				KtvAnchorEnableAddAll               string `json:"ktv_anchor_enable_add_all"`
				KtvAutoMuteSelf                     string `json:"ktv_auto_mute_self"`
				KtvChallengeMinusGift               string `json:"ktv_challenge_minus_gift"`
				KtvComponentNewMidi                 string `json:"ktv_component_new_midi"`
				KtvEnableAvatar                     string `json:"ktv_enable_avatar"`
				KtvEnableOpenCamera                 string `json:"ktv_enable_open_camera"`
				KtvFragmentSong                     string `json:"ktv_fragment_song"`
				KtvGrabGuideSong                    string `json:"ktv_grab_guide_song"`
				KtvGuideSongSwitch                  string `json:"ktv_guide_song_switch"`
				KtvKickWhenLinkerFull               string `json:"ktv_kick_when_linker_full"`
				KtvMcHostShowTag                    string `json:"ktv_mc_host_show_tag"`
				KtvNewChallenge                     string `json:"ktv_new_challenge"`
				KtvRoomAtmosphere                   string `json:"ktv_room_atmosphere"`
				KtvSingingHotRank                   string `json:"ktv_singing_hot_rank"`
				KtvVideoStreamOptimize              string `json:"ktv_video_stream_optimize"`
				KtvWantListenEnable                 string `json:"ktv_want_listen_enable"`
				LinkmicMultiChorus                  string `json:"linkmic_multi_chorus"`
				LinkmicOrderSingSearchFingerprint   string `json:"linkmic_order_sing_search_fingerprint"`
				LinkmicOrderSingUpgrade             string `json:"linkmic_order_sing_upgrade"`
				LinkmicStarwish                     string `json:"linkmic_starwish"`
				LinkmicVideoEqualLayoutFrameOpt     string `json:"linkmic_video_equal_layout_frame_opt"`
				LiveAnchorEnableChorus              string `json:"live_anchor_enable_chorus"`
				LiveAnchorEnableCustomPosition      string `json:"live_anchor_enable_custom_position"`
				LiveAnchorHitNewAudienceLinkmic     string `json:"live_anchor_hit_new_audience_linkmic"`
				LiveAnchorHitPositionOpt            string `json:"live_anchor_hit_position_opt"`
				LiveAnchorHitVideoBidPaid           string `json:"live_anchor_hit_video_bid_paid"`
				LiveAnchorHitVideoTeamfight         string `json:"live_anchor_hit_video_teamfight"`
				LiveAnswerOnWall                    string `json:"live_answer_on_wall"`
				LiveAudienceLinkmicPreApplyV2       string `json:"live_audience_linkmic_pre_apply_v2"`
				LiveAudioEnableCPosition            string `json:"live_audio_enable_c_position"`
				LiveBackupSeiEnable                 string `json:"live_backup_sei_enable"`
				LiveDouPlusEnter                    string `json:"live_dou_plus_enter"`
				LiveKtvEnableBeat                   string `json:"live_ktv_enable_beat"`
				LiveKtvGroup                        string `json:"live_ktv_group"`
				LiveKtvShowSingerIcon               string `json:"live_ktv_show_singer_icon"`
				LiveKtvSingingChallenge             string `json:"live_ktv_singing_challenge"`
				LiveLinkmicBattleOptimize           string `json:"live_linkmic_battle_optimize"`
				LiveLinkmicKtvAnchorLyricMode       string `json:"live_linkmic_ktv_anchor_lyric_mode"`
				LiveLinkmicOrderSingMicroOpt        string `json:"live_linkmic_order_sing_micro_opt"`
				LiveLinkmicOrderSingV3              string `json:"live_linkmic_order_sing_v3"`
				LivePcHelperNewLayout               string `json:"live_pc_helper_new_layout"`
				LiveRoomManageStyle                 string `json:"live_room_manage_style"`
				LiveTeamFightFlexible               string `json:"live_team_fight_flexible"`
				LiveVideoEnableCPosition            string `json:"live_video_enable_c_position"`
				LiveVideoEnableSelfDiscipline       string `json:"live_video_enable_self_discipline"`
				LiveVideoHostIdentityEnable         string `json:"live_video_host_identity_enable"`
				LiveVideoShare                      string `json:"live_video_share"`
				LonelyRoomEnterMsgUnfold            string `json:"lonely_room_enter_msg_unfold"`
				MarkUser                            string `json:"mark_user"`
				MergeKtvModeEnable                  string `json:"merge_ktv_mode_enable"`
				MergeKtvOptimizeEnable              string `json:"merge_ktv_optimize_enable"`
				MultiLinkBizAccessBackupSeiConfig   string `json:"multi_link_biz_access_backup_sei_config"`
				OptAudienceLinkmic                  string `json:"opt_audience_linkmic"`
				OptPaidLinkFeatureSwitch            string `json:"opt_paid_link_feature_switch"`
				OptranPaidLinkmic                   string `json:"optran_paid_linkmic"`
				OrderSingMv                         string `json:"order_sing_mv"`
				PlayModeOpt24                       string `json:"play_mode_opt_24"`
				PsUseNewPanel                       string `json:"ps_use_new_panel"`
				RadioPrepareApply                   string `json:"radio_prepare_apply"`
				RoomBattleModeSwitch                string `json:"room_battle_mode_switch"`
				RoomBattleModeSwitchAudio           string `json:"room_battle_mode_switch_audio"`
				RoomBattleVideoAudioInterconnection string `json:"room_battle_video_audio_interconnection"`
				RoomDoubleLike                      string `json:"room_double_like"`
				RoomSecretChat                      string `json:"room_secret_chat"`
				SelfDisciplineV2                    string `json:"self_discipline_v2"`
				SelfDisciplineV3                    string `json:"self_discipline_v3"`
				SocialShareVideoAdjustVolume        string `json:"social_share_video_adjust_volume"`
				SupportMultipleAddPrice             string `json:"support_multiple_add_price"`
				ThemedCompetitionV2                 string `json:"themed_competition_v2"`
				TrafficStrategy                     string `json:"traffic_strategy"`
				VideoEqual1V8FixSwitch              string `json:"video_equal_1v8fix_switch"`
				VideoKtvChallenge                   string `json:"video_ktv_challenge"`
				VideoTalkEnableAvatar               string `json:"video_talk_enable_avatar"`
			} `json:"AnchorABMap"`
			LikeCount      int    `json:"like_count"`
			OwnerUserIdStr string `json:"owner_user_id_str"`
			PaidLiveData   struct {
				PaidType           int  `json:"paid_type"`
				ViewRight          int  `json:"view_right"`
				Duration           int  `json:"duration"`
				Delivery           int  `json:"delivery"`
				NeedDeliveryNotice bool `json:"need_delivery_notice"`
				AnchorRight        int  `json:"anchor_right"`
				PayAbType          int  `json:"pay_ab_type"`
				PrivilegeInfo      struct {
				} `json:"privilege_info"`
				PrivilegeInfoMap struct {
				} `json:"privilege_info_map"`
				MaxPreviewDuration int `json:"max_preview_duration"`
			} `json:"paid_live_data"`
			Basis struct {
				NextPing             int  `json:"next_ping"`
				IsCustomizeAudioRoom bool `json:"is_customize_audio_room"`
				NeedRequestLuckybox  int  `json:"need_request_luckybox"`
				SecretRoom           int  `json:"secret_room"`
				ForeignUserRoom      int  `json:"foreign_user_room"`
			} `json:"basis"`
			GameData struct {
				GameTagInfo struct {
					IsGame      int    `json:"is_game"`
					GameTagId   int    `json:"game_tag_id"`
					GameTagName string `json:"game_tag_name"`
				} `json:"game_tag_info"`
			} `json:"game_data"`
			PicoInfo struct {
				PicoLiveType                 int    `json:"pico_live_type"`
				PicoVirtualLiveBgImageUri    string `json:"pico_virtual_live_bg_image_uri"`
				PicoCreateScene              string `json:"pico_create_scene"`
				CustomInfo                   string `json:"custom_info"`
				PicoVirtualLiveBgImageDigest string `json:"pico_virtual_live_bg_image_digest"`
				VirtualLiveBgImages          struct {
					OriginalDigest  string        `json:"original_digest"`
					IsUpright       bool          `json:"is_upright"`
					ConvertedImages []interface{} `json:"converted_images"`
					ConvertedList   []interface{} `json:"converted_list"`
				} `json:"virtual_live_bg_images"`
				Pitch          int `json:"pitch"`
				ClientLiveType int `json:"client_live_type"`
				PicoVrTransfer int `json:"pico_vr_transfer"`
				PicoLiveMode   int `json:"pico_live_mode"`
				StreamMapping  struct {
				} `json:"stream_mapping"`
			} `json:"pico_info"`
			ShortTouchAreaConfig struct {
				Elements struct {
					Field1 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"1"`
					Field2 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"2"`
					Field3 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"3"`
					Field4 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"4"`
					Field5 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"5"`
					Field6 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"6"`
					Field7 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"7"`
					Field8 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"8"`
					Field9 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"9"`
					Field10 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"10"`
					Field11 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"12"`
					Field12 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"22"`
					Field13 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"27"`
					Field14 struct {
						Type     int `json:"type"`
						Priority int `json:"priority"`
					} `json:"30"`
				} `json:"elements"`
				ForbiddenTypesMap struct {
				} `json:"forbidden_types_map"`
				TempStateConditionMap struct {
					Field1 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"1"`
					Field2 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"2"`
					Field3 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"3"`
					Field4 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"4"`
					Field5 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"5"`
					Field6 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"6"`
					Field7 struct {
						Type struct {
							StrategyType int `json:"strategy_type"`
							Priority     int `json:"priority"`
						} `json:"type"`
						MinimumGap int `json:"minimum_gap"`
					} `json:"7"`
				} `json:"temp_state_condition_map"`
				TempStateStrategy struct {
					Field1 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Field3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
							Field4 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"6"`
							Field5 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"7"`
						} `json:"strategy_map"`
					} `json:"4"`
					Field2 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Field3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
							Field4 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"4"`
							Field5 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"5"`
							Field6 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"6"`
						} `json:"strategy_map"`
					} `json:"7"`
					Field3 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"8"`
					Field4 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Field3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
							Field4 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"5"`
							Field5 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"6"`
							Field6 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"7"`
						} `json:"strategy_map"`
					} `json:"97"`
					Field5 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"136"`
					Field6 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Field3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
						} `json:"strategy_map"`
					} `json:"141"`
					Field7 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"149"`
					Field8 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"152"`
					Field9 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
							Field3 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"4"`
						} `json:"strategy_map"`
					} `json:"153"`
					Field10 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
						} `json:"strategy_map"`
					} `json:"159"`
					Field11 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
							Field2 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"161"`
					Field12 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
						} `json:"strategy_map"`
					} `json:"210"`
					Field13 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
						} `json:"strategy_map"`
					} `json:"217"`
					Field14 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
						} `json:"strategy_map"`
					} `json:"222"`
					Field15 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
						} `json:"strategy_map"`
					} `json:"306"`
					Field16 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"4"`
						} `json:"strategy_map"`
					} `json:"307"`
					Field17 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"5"`
						} `json:"strategy_map"`
					} `json:"308"`
					Field18 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"3"`
						} `json:"strategy_map"`
					} `json:"311"`
					Field19 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"1"`
						} `json:"strategy_map"`
					} `json:"312"`
					Field20 struct {
						ShortTouchType int `json:"short_touch_type"`
						StrategyMap    struct {
							Field1 struct {
								Type struct {
									StrategyType int `json:"strategy_type"`
									Priority     int `json:"priority"`
								} `json:"type"`
								Duration       int    `json:"duration"`
								StrategyMethod string `json:"strategy_method"`
							} `json:"2"`
						} `json:"strategy_map"`
					} `json:"313"`
				} `json:"temp_state_strategy"`
				StrategyFeatWhitelist    []string `json:"strategy_feat_whitelist"`
				TempStateGlobalCondition struct {
					DurationGap         int   `json:"duration_gap"`
					AllowCount          int   `json:"allow_count"`
					IgnoreStrategyTypes []int `json:"ignore_strategy_types"`
				} `json:"temp_state_global_condition"`
			} `json:"short_touch_area_config"`
			ReqUser struct {
				UserShareRoomScore  int `json:"user_share_room_score"`
				EnterUserDeviceType int `json:"enter_user_device_type"`
			} `json:"req_user"`
			Others struct {
				DecoDetail struct {
				} `json:"deco_detail"`
				MorePanelInfo struct {
					LoadStrategy int `json:"load_strategy"`
				} `json:"more_panel_info"`
				AppointmentInfo struct {
					AppointmentId int  `json:"appointment_id"`
					IsSubscribe   bool `json:"is_subscribe"`
				} `json:"appointment_info"`
				WebSkin struct {
					EnableSkin bool `json:"enable_skin"`
				} `json:"web_skin"`
				Programme struct {
					EnableProgramme bool `json:"enable_programme"`
				} `json:"programme"`
				WebLivePortOptimization struct {
					StrategyConfig struct {
						Background struct {
							StrategyType         int    `json:"strategy_type"`
							UseConfigDuration    bool   `json:"use_config_duration"`
							PauseMonitorDuration string `json:"pause_monitor_duration"`
						} `json:"background"`
						Detail struct {
							StrategyType         int    `json:"strategy_type"`
							UseConfigDuration    bool   `json:"use_config_duration"`
							PauseMonitorDuration string `json:"pause_monitor_duration"`
						} `json:"detail"`
						Tab struct {
							StrategyType         int    `json:"strategy_type"`
							UseConfigDuration    bool   `json:"use_config_duration"`
							PauseMonitorDuration string `json:"pause_monitor_duration"`
						} `json:"tab"`
					} `json:"strategy_config"`
					StrategyExtra string `json:"strategy_extra"`
				} `json:"web_live_port_optimization"`
				LvideoItemId          int `json:"lvideo_item_id"`
				RecognitionContainers struct {
					RecognitionCandidates []interface{} `json:"recognition_candidates"`
				} `json:"recognition_containers"`
				AnchorTogetherLive struct {
					IsTogetherLive int           `json:"is_together_live"`
					UserList       []interface{} `json:"user_list"`
					Title          string        `json:"title"`
					SchemaUrl      string        `json:"schema_url"`
					Scene          int           `json:"scene"`
					IsShow         bool          `json:"is_show"`
				} `json:"anchor_together_live"`
				MosaicVersion         int           `json:"mosaic_version"`
				MetricTrackerDataList []interface{} `json:"metric_tracker_data_list"`
				CloudCollaborateData  struct {
					CollaborateRoomId    int    `json:"collaborate_room_id"`
					CollaborateRoomIdStr string `json:"collaborate_room_id_str"`
					RejoinTime           int    `json:"rejoin_time"`
				} `json:"cloud_collaborate_data"`
				RoomChatGuideLocaleCity string `json:"room_chat_guide_locale_city"`
			} `json:"others"`
			AdminUserOpenIds    []interface{} `json:"admin_user_open_ids"`
			AdminUserOpenIdsStr []interface{} `json:"admin_user_open_ids_str"`
			OwnerOpenIdStr      string        `json:"owner_open_id_str"`
		} `json:"data"`
		EnterRoomId string `json:"enter_room_id"`
		Extra       struct {
			DiggColor         string `json:"digg_color"`
			PayScores         string `json:"pay_scores"`
			IsOfficialChannel bool   `json:"is_official_channel"`
			Signature         string `json:"signature"`
			VrType            int    `json:"vr_type"`
		} `json:"extra"`
		User struct {
			IdStr       string `json:"id_str"`
			SecUid      string `json:"sec_uid"`
			Nickname    string `json:"nickname"`
			AvatarThumb struct {
				UrlList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			FollowInfo struct {
				FollowStatus    int    `json:"follow_status"`
				FollowStatusStr string `json:"follow_status_str"`
			} `json:"follow_info"`
			ForeignUser int    `json:"foreign_user"`
			OpenIdStr   string `json:"open_id_str"`
		} `json:"user"`
		QrcodeUrl        string `json:"qrcode_url"`
		EnterMode        int    `json:"enter_mode"`
		RoomStatus       int    `json:"room_status"`
		PartitionRoadMap struct {
		} `json:"partition_road_map"`
		SimilarRooms      []interface{} `json:"similar_rooms"`
		SharkDecisionConf string        `json:"shark_decision_conf"`
		WebStreamUrl      struct {
			FlvPullUrl struct {
			} `json:"flv_pull_url"`
			DefaultResolution string `json:"default_resolution"`
			HlsPullUrlMap     struct {
			} `json:"hls_pull_url_map"`
			HlsPullUrl        string `json:"hls_pull_url"`
			StreamOrientation int    `json:"stream_orientation"`
			PullDatas         struct {
			} `json:"pull_datas"`
		} `json:"web_stream_url"`
		LoginLead struct {
			IsLogin bool `json:"is_login"`
			Level   int  `json:"level"`
			Items   struct {
			} `json:"items"`
		} `json:"login_lead"`
		AuthCertInfo string `json:"auth_cert_info"`
	} `json:"data"`
	Extra struct {
		Now int64 `json:"now"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}
