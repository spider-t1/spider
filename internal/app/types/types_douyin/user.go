package types_douyin

import "spider/internal/app/types/types_common"

// UserInfoReq 抖音用户信息请求参数
type UserInfoReq struct {
	types_common.BaseListParam
	SecUserId string `form:"sec_user_id" binding:"required"`
}

func (r *UserInfoReq) Adjust() {}

// UserSearchReq 抖音用户搜索请求参数
type UserSearchReq struct {
	types_common.BaseListParam
	Keyword string `form:"keyword" binding:"required"`
}

func (r *UserSearchReq) Adjust() {}

// UserVideoReq 抖音用户视频列表请求参数
type UserVideoReq struct {
	types_common.BaseListParam
	SecId string `form:"sec_id" binding:"required"`
}

func (r *UserVideoReq) Adjust() {}

type DouyinSearchResp struct {
	Type          int          `json:"type"`
	UserList      []DouyinUser `json:"user_list"`
	ChallengeList interface{}  `json:"challenge_list"`
	MusicList     interface{}  `json:"music_list"`
	Cursor        int          `json:"cursor"`
	HasMore       int          `json:"has_more"`
	StatusCode    int          `json:"status_code"`
	Qc            string       `json:"qc"`
	MyselfUserId  string       `json:"myself_user_id"`
	Rid           string       `json:"rid"`
	LogPb         struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	Extra struct {
		Now             int64         `json:"now"`
		Logid           string        `json:"logid"`
		FatalItemIds    []interface{} `json:"fatal_item_ids"`
		SearchRequestId string        `json:"search_request_id"`
		Scenes          interface{}   `json:"scenes"`
	} `json:"extra"`
	InputKeyword       string `json:"input_keyword"`
	GlobalDoodleConfig struct {
		Keyword        string `json:"keyword"`
		FilterShowDot  int    `json:"filter_show_dot"`
		FilterSettings []struct {
			Items []struct {
				Title    string `json:"title"`
				Value    string `json:"value"`
				LogValue string `json:"log_value"`
			} `json:"items"`
			Title        string `json:"title"`
			Name         string `json:"name"`
			DefaultIndex int    `json:"default_index"`
			LogName      string `json:"log_name"`
		} `json:"filter_settings"`
	} `json:"global_doodle_config"`
	Path string `json:"path"`
}

type DouyinUser struct {
	UserInfo struct {
		Uid         string `json:"uid"`
		ShortId     string `json:"short_id"`
		Nickname    string `json:"nickname"`
		Signature   string `json:"signature"`
		AvatarThumb struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
			Height  int      `json:"height"`
		} `json:"avatar_thumb"`
		FollowStatus           int         `json:"follow_status"`
		FollowerCount          int         `json:"follower_count"`
		TotalFavorited         int         `json:"total_favorited"`
		CustomVerify           string      `json:"custom_verify"`
		UniqueId               string      `json:"unique_id"`
		RoomId                 int         `json:"room_id,omitempty"`
		EnterpriseVerifyReason string      `json:"enterprise_verify_reason"`
		FollowersDetail        interface{} `json:"followers_detail"`
		PlatformSyncInfo       interface{} `json:"platform_sync_info"`
		Secret                 int         `json:"secret"`
		Geofencing             interface{} `json:"geofencing"`
		FollowerStatus         int         `json:"follower_status"`
		CoverUrl               interface{} `json:"cover_url"`
		ItemList               interface{} `json:"item_list"`
		NewStoryCover          interface{} `json:"new_story_cover"`
		TypeLabel              interface{} `json:"type_label"`
		AdCoverUrl             interface{} `json:"ad_cover_url"`
		RelativeUsers          interface{} `json:"relative_users"`
		ChaList                interface{} `json:"cha_list"`
		SecUid                 string      `json:"sec_uid"`
		NeedPoints             interface{} `json:"need_points"`
		HomepageBottomToast    interface{} `json:"homepage_bottom_toast"`
		RoomData               string      `json:"room_data"`
		CanSetGeofencing       interface{} `json:"can_set_geofencing"`
		RoomIdStr              string      `json:"room_id_str,omitempty"`
		WhiteCoverUrl          interface{} `json:"white_cover_url"`
		UserTags               []struct {
			Description string `json:"description"`
			IconUrl     string `json:"icon_url"`
			Type        string `json:"type"`
		} `json:"user_tags"`
		BanUserFunctions                       interface{} `json:"ban_user_functions"`
		VersatileDisplay                       string      `json:"versatile_display"`
		CardEntries                            interface{} `json:"card_entries"`
		DisplayInfo                            interface{} `json:"display_info"`
		CardEntriesNotDisplay                  interface{} `json:"card_entries_not_display"`
		CardSortPriority                       interface{} `json:"card_sort_priority"`
		InterestTags                           interface{} `json:"interest_tags"`
		LinkItemList                           interface{} `json:"link_item_list"`
		UserPermissions                        interface{} `json:"user_permissions"`
		OfflineInfoList                        interface{} `json:"offline_info_list"`
		SignatureExtra                         interface{} `json:"signature_extra"`
		PersonalTagList                        interface{} `json:"personal_tag_list"`
		CfList                                 interface{} `json:"cf_list"`
		ImRoleIds                              interface{} `json:"im_role_ids"`
		NotSeenItemIdList                      interface{} `json:"not_seen_item_id_list"`
		FollowerListSecondaryInformationStruct interface{} `json:"follower_list_secondary_information_struct"`
		EndorsementInfoList                    interface{} `json:"endorsement_info_list"`
		TextExtra                              interface{} `json:"text_extra"`
		ContrailList                           interface{} `json:"contrail_list"`
		DataLabelList                          interface{} `json:"data_label_list"`
		NotSeenItemIdListV2                    interface{} `json:"not_seen_item_id_list_v2"`
		SpecialPeopleLabels                    interface{} `json:"special_people_labels"`
		FamiliarVisitorUser                    interface{} `json:"familiar_visitor_user"`
		AvatarSchemaList                       interface{} `json:"avatar_schema_list"`
		ProfileMobParams                       interface{} `json:"profile_mob_params"`
		VerificationPermissionIds              interface{} `json:"verification_permission_ids"`
		BatchUnfollowRelationDesc              interface{} `json:"batch_unfollow_relation_desc"`
		BatchUnfollowContainTabs               interface{} `json:"batch_unfollow_contain_tabs"`
		CreatorTagList                         interface{} `json:"creator_tag_list"`
		AccountCertInfo                        string      `json:"account_cert_info,omitempty"`
		PrivateRelationList                    interface{} `json:"private_relation_list"`
		FollowerCountStr                       string      `json:"follower_count_str"`
		IdentityLabels                         interface{} `json:"identity_labels"`
		ProfileComponentDisabled               interface{} `json:"profile_component_disabled"`
	} `json:"user_info"`
	Position        interface{} `json:"position"`
	UniqidPosition  interface{} `json:"uniqid_position"`
	Effects         interface{} `json:"effects"`
	Musics          interface{} `json:"musics"`
	Items           interface{} `json:"items"`
	MixList         interface{} `json:"mix_list"`
	Challenges      interface{} `json:"challenges"`
	ProductInfo     interface{} `json:"product_info"`
	ProductList     interface{} `json:"product_list"`
	IsRedUniqueid   bool        `json:"is_red_uniqueid"`
	Baikes          interface{} `json:"baikes"`
	UserSubLightApp interface{} `json:"userSubLightApp"`
	ShopProductInfo interface{} `json:"shop_product_info"`
	UserServiceInfo interface{} `json:"user_service_info"`
	Fandoms         interface{} `json:"fandoms"`
}
