package types_douyin

type DyOneVideoInfo struct {
	StatusCode int   `json:"status_code"`
	MinCursor  int   `json:"min_cursor"`
	MaxCursor  int64 `json:"max_cursor"`
	HasMore    int   `json:"has_more"`
	AwemeList  []struct {
		AwemeId    string `json:"aweme_id"`
		Desc       string `json:"desc"`
		CreateTime int    `json:"create_time"`
		Author     struct {
			Uid            string      `json:"uid"`
			SignatureExtra interface{} `json:"signature_extra"`
			Nickname       string      `json:"nickname"`
			DisplayInfo    interface{} `json:"display_info"`
			NeedPoints     interface{} `json:"need_points"`
			AvatarThumb    struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_thumb"`
			BatchUnfollowRelationDesc interface{} `json:"batch_unfollow_relation_desc"`
			ImRoleIds                 interface{} `json:"im_role_ids"`
			FollowStatus              int         `json:"follow_status"`
			TextExtra                 interface{} `json:"text_extra"`
			StoryInteractive          int         `json:"story_interactive"`
			ContrailList              interface{} `json:"contrail_list"`
			CustomVerify              string      `json:"custom_verify"`
			VerificationPermissionIds interface{} `json:"verification_permission_ids"`
			CardEntries               interface{} `json:"card_entries"`
			CardEntriesNotDisplay     interface{} `json:"card_entries_not_display"`
			CreatorTagList            interface{} `json:"creator_tag_list"`
			WhiteCoverUrl             interface{} `json:"white_cover_url"`
			ShareInfo                 struct {
				ShareUrl       string `json:"share_url"`
				ShareWeiboDesc string `json:"share_weibo_desc"`
				ShareDesc      string `json:"share_desc"`
				ShareTitle     string `json:"share_title"`
				ShareQrcodeUrl struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
					Width   int      `json:"width"`
					Height  int      `json:"height"`
				} `json:"share_qrcode_url"`
				ShareTitleMyself string `json:"share_title_myself"`
				ShareTitleOther  string `json:"share_title_other"`
				ShareDescInfo    string `json:"share_desc_info"`
			} `json:"share_info"`
			HomepageBottomToast                    interface{} `json:"homepage_bottom_toast"`
			UserTags                               interface{} `json:"user_tags"`
			SecUid                                 string      `json:"sec_uid"`
			EnterpriseVerifyReason                 string      `json:"enterprise_verify_reason"`
			IsAdFake                               bool        `json:"is_ad_fake"`
			StoryTtl                               int         `json:"story_ttl"`
			InterestTags                           interface{} `json:"interest_tags"`
			BanUserFunctions                       interface{} `json:"ban_user_functions"`
			BatchUnfollowContainTabs               interface{} `json:"batch_unfollow_contain_tabs"`
			LinkItemList                           interface{} `json:"link_item_list"`
			EndorsementInfoList                    interface{} `json:"endorsement_info_list"`
			OfflineInfoList                        interface{} `json:"offline_info_list"`
			PreventDownload                        bool        `json:"prevent_download"`
			NotSeenItemIdListV2                    interface{} `json:"not_seen_item_id_list_v2"`
			CanSetGeofencing                       interface{} `json:"can_set_geofencing"`
			FollowerListSecondaryInformationStruct interface{} `json:"follower_list_secondary_information_struct"`
			FollowerStatus                         int         `json:"follower_status"`
			AvatarSchemaList                       interface{} `json:"avatar_schema_list"`
			SpecialPeopleLabels                    interface{} `json:"special_people_labels"`
			FamiliarVisitorUser                    interface{} `json:"familiar_visitor_user"`
			CoverUrl                               []struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover_url"`
			PersonalTagList          interface{} `json:"personal_tag_list"`
			CfList                   interface{} `json:"cf_list"`
			DataLabelList            interface{} `json:"data_label_list"`
			UserPermissions          interface{} `json:"user_permissions"`
			NotSeenItemIdList        interface{} `json:"not_seen_item_id_list"`
			RiskNoticeText           string      `json:"risk_notice_text"`
			ProfileComponentDisabled interface{} `json:"profile_component_disabled"`
			ProfileMobParams         interface{} `json:"profile_mob_params"`
			IdentityLabels           interface{} `json:"identity_labels"`
			PrivateRelationList      interface{} `json:"private_relation_list"`
			CardSortPriority         interface{} `json:"card_sort_priority"`
		} `json:"author"`
		Music struct {
			Id      int64  `json:"id"`
			IdStr   string `json:"id_str"`
			Title   string `json:"title"`
			Author  string `json:"author"`
			Album   string `json:"album"`
			CoverHd struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover_hd"`
			CoverLarge struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover_large"`
			CoverMedium struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover_medium"`
			CoverThumb struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover_thumb"`
			PlayUrl struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
				UrlKey  string   `json:"url_key"`
			} `json:"play_url"`
			SchemaUrl         string      `json:"schema_url"`
			SourcePlatform    int         `json:"source_platform"`
			StartTime         int         `json:"start_time"`
			EndTime           int         `json:"end_time"`
			Duration          int         `json:"duration"`
			Extra             string      `json:"extra"`
			UserCount         int         `json:"user_count"`
			Position          interface{} `json:"position"`
			CollectStat       int         `json:"collect_stat"`
			Status            int         `json:"status"`
			OfflineDesc       string      `json:"offline_desc"`
			OwnerId           string      `json:"owner_id,omitempty"`
			OwnerNickname     string      `json:"owner_nickname"`
			IsOriginal        bool        `json:"is_original"`
			Mid               string      `json:"mid"`
			BindedChallengeId int         `json:"binded_challenge_id"`
			Redirect          bool        `json:"redirect"`
			IsRestricted      bool        `json:"is_restricted"`
			AuthorDeleted     bool        `json:"author_deleted"`
			IsDelVideo        bool        `json:"is_del_video"`
			IsVideoSelfSee    bool        `json:"is_video_self_see"`
			OwnerHandle       string      `json:"owner_handle"`
			AuthorPosition    interface{} `json:"author_position"`
			PreventDownload   bool        `json:"prevent_download"`
			StrongBeatUrl     struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"strong_beat_url,omitempty"`
			UnshelveCountries         interface{}   `json:"unshelve_countries"`
			PreventItemDownloadStatus int           `json:"prevent_item_download_status"`
			ExternalSongInfo          []interface{} `json:"external_song_info"`
			SecUid                    string        `json:"sec_uid,omitempty"`
			AvatarThumb               struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_thumb,omitempty"`
			AvatarMedium struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_medium,omitempty"`
			AvatarLarge struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"avatar_large,omitempty"`
			PreviewStartTime int  `json:"preview_start_time"`
			PreviewEndTime   int  `json:"preview_end_time"`
			IsCommerceMusic  bool `json:"is_commerce_music"`
			IsOriginalSound  bool `json:"is_original_sound"`
			AuditionDuration int  `json:"audition_duration"`
			ShootDuration    int  `json:"shoot_duration"`
			ReasonType       int  `json:"reason_type"`
			Artists          []struct {
				Uid      string `json:"uid"`
				SecUid   string `json:"sec_uid"`
				NickName string `json:"nick_name"`
				Handle   string `json:"handle"`
				Avatar   struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"avatar"`
				IsVerified bool `json:"is_verified"`
				EnterType  int  `json:"enter_type"`
			} `json:"artists"`
			LyricShortPosition   interface{} `json:"lyric_short_position"`
			MuteShare            bool        `json:"mute_share"`
			TagList              interface{} `json:"tag_list"`
			DmvAutoShow          bool        `json:"dmv_auto_show"`
			IsPgc                bool        `json:"is_pgc"`
			IsMatchedMetadata    bool        `json:"is_matched_metadata"`
			IsAudioUrlWithCookie bool        `json:"is_audio_url_with_cookie"`
			MusicChartRanks      interface{} `json:"music_chart_ranks"`
			CanBackgroundPlay    bool        `json:"can_background_play"`
			MusicStatus          int         `json:"music_status"`
			VideoDuration        int         `json:"video_duration"`
			PgcMusicType         int         `json:"pgc_music_type"`
			AuthorStatus         int         `json:"author_status,omitempty"`
			SearchImpr           struct {
				EntityId string `json:"entity_id"`
			} `json:"search_impr"`
			ArtistUserInfos                interface{} `json:"artist_user_infos"`
			DspStatus                      int         `json:"dsp_status"`
			MusicianUserInfos              interface{} `json:"musician_user_infos"`
			MusicCollectCount              int         `json:"music_collect_count"`
			MusicCoverAtmosphereColorValue string      `json:"music_cover_atmosphere_color_value"`
			ShowOriginClip                 bool        `json:"show_origin_clip"`
			CoverColorHsv                  struct {
				H int `json:"h"`
				S int `json:"s"`
				V int `json:"v"`
			} `json:"cover_color_hsv,omitempty"`
			Song struct {
				Id      int64       `json:"id"`
				IdStr   string      `json:"id_str"`
				Title   string      `json:"title,omitempty"`
				Artists interface{} `json:"artists"`
				Chorus  struct {
					StartMs    int `json:"start_ms"`
					DurationMs int `json:"duration_ms"`
				} `json:"chorus,omitempty"`
				ChorusV3Infos interface{} `json:"chorus_v3_infos"`
			} `json:"song,omitempty"`
			MusicImageBeats struct {
				MusicImageBeatsUrl struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
					Width   int      `json:"width"`
					Height  int      `json:"height"`
				} `json:"music_image_beats_url"`
				MusicImageBeatsRaw string `json:"music_image_beats_raw,omitempty"`
			} `json:"music_image_beats,omitempty"`
			MatchedPgcSound struct {
				Author      string `json:"author"`
				Title       string `json:"title"`
				MixedTitle  string `json:"mixed_title"`
				MixedAuthor string `json:"mixed_author"`
				CoverMedium struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
					Width   int      `json:"width"`
					Height  int      `json:"height"`
				} `json:"cover_medium"`
			} `json:"matched_pgc_sound,omitempty"`
		} `json:"music"`
		ProductGenreInfo struct {
			ProductGenreType        int   `json:"product_genre_type"`
			MaterialGenreSubTypeSet []int `json:"material_genre_sub_type_set"`
			SpecialInfo             struct {
				RecommendGroupName int `json:"recommend_group_name"`
			} `json:"special_info"`
		} `json:"product_genre_info"`
		Video struct {
			PlayAddr struct {
				Uri      string   `json:"uri"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
				Height   int      `json:"height"`
				UrlKey   string   `json:"url_key"`
				DataSize int      `json:"data_size,omitempty"`
				FileHash string   `json:"file_hash,omitempty"`
				FileCs   string   `json:"file_cs,omitempty"`
			} `json:"play_addr"`
			Cover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"cover"`
			Height       int `json:"height"`
			Width        int `json:"width"`
			DynamicCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"dynamic_cover,omitempty"`
			OriginCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"origin_cover"`
			Ratio       string `json:"ratio"`
			Format      string `json:"format,omitempty"`
			Meta        string `json:"meta"`
			IsSourceHDR int    `json:"is_source_HDR,omitempty"`
			BitRate     []struct {
				GearName    string `json:"gear_name"`
				QualityType int    `json:"quality_type"`
				BitRate     int    `json:"bit_rate"`
				PlayAddr    struct {
					Uri      string   `json:"uri"`
					UrlList  []string `json:"url_list"`
					Width    int      `json:"width"`
					Height   int      `json:"height"`
					UrlKey   string   `json:"url_key"`
					DataSize int      `json:"data_size"`
					FileHash string   `json:"file_hash"`
					FileCs   string   `json:"file_cs"`
				} `json:"play_addr"`
				IsH265     int    `json:"is_h265"`
				IsBytevc1  int    `json:"is_bytevc1"`
				HDRType    string `json:"HDR_type"`
				HDRBit     string `json:"HDR_bit"`
				FPS        int    `json:"FPS"`
				VideoExtra string `json:"video_extra"`
				Format     string `json:"format"`
			} `json:"bit_rate"`
			Duration      int         `json:"duration"`
			BitRateAudio  interface{} `json:"bit_rate_audio"`
			GaussianCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"gaussian_cover,omitempty"`
			PlayAddr265 struct {
				Uri      string   `json:"uri"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
				Height   int      `json:"height"`
				UrlKey   string   `json:"url_key"`
				DataSize int      `json:"data_size"`
				FileHash string   `json:"file_hash"`
				FileCs   string   `json:"file_cs"`
			} `json:"play_addr_265,omitempty"`
			Audio struct {
				OriginalSoundInfos interface{} `json:"original_sound_infos"`
			} `json:"audio"`
			PlayAddrH264 struct {
				Uri      string   `json:"uri"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
				Height   int      `json:"height"`
				UrlKey   string   `json:"url_key"`
				DataSize int      `json:"data_size"`
				FileHash string   `json:"file_hash"`
				FileCs   string   `json:"file_cs"`
			} `json:"play_addr_h264,omitempty"`
			RawCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"raw_cover,omitempty"`
			AnimatedCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"animated_cover,omitempty"`
			HorizontalType int `json:"horizontal_type,omitempty"`
			BigThumbs      []struct {
				ImgNum   int      `json:"img_num"`
				Uri      string   `json:"uri"`
				ImgUrl   string   `json:"img_url"`
				ImgXSize int      `json:"img_x_size"`
				ImgYSize int      `json:"img_y_size"`
				ImgXLen  int      `json:"img_x_len"`
				ImgYLen  int      `json:"img_y_len"`
				Duration float64  `json:"duration"`
				Interval int      `json:"interval"`
				Fext     string   `json:"fext"`
				Uris     []string `json:"uris"`
				ImgUrls  []string `json:"img_urls"`
			} `json:"big_thumbs"`
			VideoModel     string `json:"video_model,omitempty"`
			UseStaticCover bool   `json:"use_static_cover,omitempty"`
		} `json:"video"`
		ShareUrl   string `json:"share_url"`
		UserDigged int    `json:"user_digged"`
		Statistics struct {
			RecommendCount int `json:"recommend_count"`
			CommentCount   int `json:"comment_count"`
			DiggCount      int `json:"digg_count"`
			AdmireCount    int `json:"admire_count"`
			PlayCount      int `json:"play_count"`
			ShareCount     int `json:"share_count"`
			CollectCount   int `json:"collect_count"`
		} `json:"statistics"`
		Status struct {
			NotAllowSoftDelReason string `json:"not_allow_soft_del_reason"`
			IsDelete              bool   `json:"is_delete"`
			AllowShare            bool   `json:"allow_share"`
			ReviewResult          struct {
				ReviewStatus int `json:"review_status"`
			} `json:"review_result"`
			AllowFriendRecommendGuide  bool `json:"allow_friend_recommend_guide"`
			PartSee                    int  `json:"part_see"`
			PrivateStatus              int  `json:"private_status"`
			ListenVideoStatus          int  `json:"listen_video_status"`
			InReviewing                bool `json:"in_reviewing"`
			AllowSelfRecommendToFriend bool `json:"allow_self_recommend_to_friend"`
			AllowFriendRecommend       bool `json:"allow_friend_recommend"`
			IsProhibited               bool `json:"is_prohibited"`
			EnableSoftDelete           int  `json:"enable_soft_delete"`
		} `json:"status"`
		PersonalPageBottonDiagnoseStyle int `json:"personal_page_botton_diagnose_style"`
		TextExtra                       []struct {
			Start        int    `json:"start"`
			End          int    `json:"end"`
			Type         int    `json:"type"`
			HashtagName  string `json:"hashtag_name,omitempty"`
			HashtagId    string `json:"hashtag_id,omitempty"`
			IsCommerce   bool   `json:"is_commerce,omitempty"`
			CaptionStart int    `json:"caption_start"`
			CaptionEnd   int    `json:"caption_end"`
			UserId       string `json:"user_id,omitempty"`
			SecUid       string `json:"sec_uid,omitempty"`
			SubType      int    `json:"sub_type,omitempty"`
			LiveData     string `json:"live_data,omitempty"`
		} `json:"text_extra"`
		IsTop       int `json:"is_top"`
		GameTagInfo struct {
			IsGame bool `json:"is_game"`
		} `json:"game_tag_info"`
		ShareInfo struct {
			ShareUrl      string `json:"share_url"`
			ShareLinkDesc string `json:"share_link_desc"`
		} `json:"share_info"`
		FollowShootClipInfo struct {
			ClipVideoAll     int64 `json:"clip_video_all"`
			ClipFromPlatform int64 `json:"clip_from_platform"`
			ClipFromUser     int64 `json:"clip_from_user"`
			OriginClipId     int64 `json:"origin_clip_id"`
		} `json:"follow_shoot_clip_info,omitempty"`
		VideoLabels               interface{} `json:"video_labels"`
		EntertainmentVideoPaidWay struct {
			PaidWays            []interface{} `json:"paid_ways"`
			PaidType            int           `json:"paid_type"`
			EnableUseNewEntData bool          `json:"enable_use_new_ent_data"`
		} `json:"entertainment_video_paid_way"`
		IsAds            bool        `json:"is_ads"`
		Duration         int         `json:"duration"`
		AwemeType        int         `json:"aweme_type"`
		InterestPoints   interface{} `json:"interest_points"`
		FollowShotAssets interface{} `json:"follow_shot_assets"`
		ImageInfos       interface{} `json:"image_infos"`
		RiskInfos        struct {
			Vote     bool   `json:"vote"`
			Warn     bool   `json:"warn"`
			RiskSink bool   `json:"risk_sink"`
			Type     int    `json:"type"`
			Content  string `json:"content"`
		} `json:"risk_infos"`
		IsMomentStory int `json:"is_moment_story"`
		EntLogExtra   struct {
			LogExtra string `json:"log_extra"`
		} `json:"ent_log_extra"`
		Position                interface{}   `json:"position"`
		UniqidPosition          interface{}   `json:"uniqid_position"`
		CommentList             interface{}   `json:"comment_list"`
		AuthorUserId            int64         `json:"author_user_id"`
		ItemAigcFollowShot      int           `json:"item_aigc_follow_shot"`
		Geofencing              []interface{} `json:"geofencing"`
		IsNewTextMode           int           `json:"is_new_text_mode"`
		CfAssetsType            int           `json:"cf_assets_type"`
		Region                  string        `json:"region"`
		VideoText               interface{}   `json:"video_text"`
		TrendsInfos             interface{}   `json:"trends_infos"`
		CollectStat             int           `json:"collect_stat"`
		LabelTopText            interface{}   `json:"label_top_text"`
		Promotions              []interface{} `json:"promotions"`
		GroupId                 string        `json:"group_id"`
		PreventDownload         bool          `json:"prevent_download"`
		NicknamePosition        interface{}   `json:"nickname_position"`
		ChallengePosition       interface{}   `json:"challenge_position"`
		EntertainmentVideoType  int           `json:"entertainment_video_type"`
		Is24Story               int           `json:"is_24_story"`
		EnableCommentStickerRec bool          `json:"enable_comment_sticker_rec"`
		LongVideo               interface{}   `json:"long_video"`
		ShootWay                string        `json:"shoot_way"`
		TrendsEventTrack        string        `json:"trends_event_track"`
		VideoShareEditStatus    int           `json:"video_share_edit_status"`
		Is25Story               int           `json:"is_25_story"`
		InteractionStickers     []struct {
			Type            int         `json:"type"`
			Index           int         `json:"index"`
			TrackInfo       string      `json:"track_info"`
			Attr            string      `json:"attr"`
			TextInfo        string      `json:"text_info"`
			TextInteraction interface{} `json:"text_interaction"`
			ImageIndex      int         `json:"image_index"`
			StickerConfig   string      `json:"sticker_config"`
			FlashMobInfo    struct {
				Id              string `json:"id"`
				CreatorNickName string `json:"creator_nick_name"`
				CreatorUid      int64  `json:"creator_uid"`
				UserAvatarList  []struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
					Width   int      `json:"width"`
					Height  int      `json:"height"`
				} `json:"user_avatar_list"`
				Count              int    `json:"count"`
				CountStr           string `json:"count_str"`
				Text               string `json:"text"`
				IsCommerceFlashmob bool   `json:"is_commerce_flashmob"`
				UserHasJoined      bool   `json:"user_has_joined"`
				TypeExtra          string `json:"type_extra"`
			} `json:"flash_mob_info"`
		} `json:"interaction_stickers"`
		PublishPlusAlienation struct {
			AlienationType int `json:"alienation_type"`
		} `json:"publish_plus_alienation"`
		OriginCommentIds   interface{} `json:"origin_comment_ids"`
		CommerceConfigData interface{} `json:"commerce_config_data"`
		NearbyHotComment   interface{} `json:"nearby_hot_comment"`
		VideoControl       struct {
			AllowDownload            bool `json:"allow_download"`
			ShareType                int  `json:"share_type"`
			ShowProgressBar          int  `json:"show_progress_bar"`
			DraftProgressBar         int  `json:"draft_progress_bar"`
			AllowDuet                bool `json:"allow_duet"`
			AllowReact               bool `json:"allow_react"`
			PreventDownloadType      int  `json:"prevent_download_type"`
			AllowDynamicWallpaper    bool `json:"allow_dynamic_wallpaper"`
			TimerStatus              int  `json:"timer_status"`
			AllowMusic               bool `json:"allow_music"`
			AllowStitch              bool `json:"allow_stitch"`
			AllowDouplus             bool `json:"allow_douplus"`
			AllowShare               bool `json:"allow_share"`
			ShareGrayed              bool `json:"share_grayed"`
			DownloadIgnoreVisibility bool `json:"download_ignore_visibility"`
			DuetIgnoreVisibility     bool `json:"duet_ignore_visibility"`
			ShareIgnoreVisibility    bool `json:"share_ignore_visibility"`
			DownloadInfo             struct {
				Level    int `json:"level"`
				FailInfo struct {
					Code   int    `json:"code"`
					Reason string `json:"reason"`
					Msg    string `json:"msg"`
				} `json:"fail_info,omitempty"`
			} `json:"download_info"`
			DuetInfo struct {
				Level    int `json:"level"`
				FailInfo struct {
					Code   int    `json:"code"`
					Reason string `json:"reason"`
					Msg    string `json:"msg,omitempty"`
				} `json:"fail_info,omitempty"`
			} `json:"duet_info"`
			AllowRecord         bool   `json:"allow_record"`
			DisableRecordReason string `json:"disable_record_reason"`
			TimerInfo           struct {
			} `json:"timer_info"`
		} `json:"video_control"`
		AwemeControl struct {
			CanForward     bool `json:"can_forward"`
			CanShare       bool `json:"can_share"`
			CanComment     bool `json:"can_comment"`
			CanShowComment bool `json:"can_show_comment"`
		} `json:"aweme_control"`
		CaptionTemplateId int         `json:"caption_template_id"`
		AiFollowImages    interface{} `json:"ai_follow_images"`
		CanCacheToLocal   bool        `json:"can_cache_to_local"`
		IsMomentHistory   int         `json:"is_moment_history"`
		Anchors           interface{} `json:"anchors"`
		HybridLabel       interface{} `json:"hybrid_label"`
		GeofencingRegions interface{} `json:"geofencing_regions"`
		DouplusUserType   int         `json:"douplus_user_type"`
		AwemeAcl          struct {
			DownloadMaskPanel struct {
				Code     int `json:"code"`
				ShowType int `json:"show_type"`
			} `json:"download_mask_panel"`
		} `json:"aweme_acl,omitempty"`
		AwemeTypeTags    string `json:"aweme_type_tags"`
		IsStory          int    `json:"is_story"`
		ComponentControl struct {
			DataSourceUrl string `json:"data_source_url"`
		} `json:"component_control"`
		FlashMobTrends              int         `json:"flash_mob_trends"`
		EffectInflowEffects         interface{} `json:"effect_inflow_effects"`
		CoverLabels                 interface{} `json:"cover_labels"`
		SelectAnchorExpandedContent int         `json:"select_anchor_expanded_content"`
		MvInfo                      []struct {
			Id      string `json:"id"`
			IconUrl struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"icon_url"`
			MvType           int      `json:"mv_type"`
			Name             string   `json:"name"`
			LokiId           string   `json:"loki_id"`
			IsImageTemplate  bool     `json:"is_image_template"`
			EnableAlienation bool     `json:"enable_alienation"`
			Provider         int      `json:"provider"`
			Features         []string `json:"features"`
		} `json:"mv_info"`
		GuideBtnType    int `json:"guide_btn_type"`
		SeriesBasicInfo struct {
		} `json:"series_basic_info"`
		OriginDuetResourceUri string `json:"origin_duet_resource_uri"`
		Images                []struct {
			Uri                          string      `json:"uri"`
			UrlList                      []string    `json:"url_list"`
			DownloadUrlList              []string    `json:"download_url_list"`
			Height                       int         `json:"height"`
			Width                        int         `json:"width"`
			MaskUrlList                  interface{} `json:"mask_url_list"`
			InteractionStickers          interface{} `json:"interaction_stickers"`
			WatermarkFreeDownloadUrlList interface{} `json:"watermark_free_download_url_list"`
		} `json:"images"`
		RelationLabels interface{} `json:"relation_labels"`
		HorizontalType int         `json:"horizontal_type,omitempty"`
		IsFromAdAuth   bool        `json:"is_from_ad_auth"`
		ImpressionData struct {
			GroupIdListA   []interface{} `json:"group_id_list_a"`
			GroupIdListB   []interface{} `json:"group_id_list_b"`
			SimilarIdListA interface{}   `json:"similar_id_list_a"`
			SimilarIdListB interface{}   `json:"similar_id_list_b"`
			GroupIdListC   []interface{} `json:"group_id_list_c"`
			GroupIdListD   []interface{} `json:"group_id_list_d"`
		} `json:"impression_data"`
		AwemeListenStruct struct {
			TraceInfo string `json:"trace_info"`
		} `json:"aweme_listen_struct"`
		ChapterBarColor  interface{} `json:"chapter_bar_color"`
		LibfinsertTaskId string      `json:"libfinsert_task_id"`
		SocialTagList    interface{} `json:"social_tag_list"`
		SuggestWords     struct {
			SuggestWords []struct {
				Words     []interface{} `json:"words"`
				Scene     string        `json:"scene"`
				IconUrl   string        `json:"icon_url"`
				HintText  string        `json:"hint_text"`
				ExtraInfo string        `json:"extra_info"`
			} `json:"suggest_words"`
		} `json:"suggest_words,omitempty"`
		ShowFollowButton struct {
		} `json:"show_follow_button"`
		DuetAggregateInMusicTab bool `json:"duet_aggregate_in_music_tab"`
		IsDuetSing              bool `json:"is_duet_sing"`
		CommentPermissionInfo   struct {
			CommentPermissionStatus int  `json:"comment_permission_status"`
			CanComment              bool `json:"can_comment"`
			ItemDetailEntry         bool `json:"item_detail_entry"`
			PressEntry              bool `json:"press_entry"`
			ToastGuide              bool `json:"toast_guide"`
		} `json:"comment_permission_info"`
		OriginalImages interface{} `json:"original_images"`
		SeriesPaidInfo struct {
			SeriesPaidStatus int `json:"series_paid_status"`
			ItemPrice        int `json:"item_price"`
		} `json:"series_paid_info"`
		ImgBitrate          []interface{} `json:"img_bitrate"`
		CommentGid          int64         `json:"comment_gid"`
		ImageAlbumMusicInfo struct {
			BeginTime int `json:"begin_time"`
			EndTime   int `json:"end_time"`
			Volume    int `json:"volume"`
		} `json:"image_album_music_info"`
		VideoTag []struct {
			TagId   int    `json:"tag_id"`
			TagName string `json:"tag_name"`
			Level   int    `json:"level"`
		} `json:"video_tag"`
		IsCollectsSelected int         `json:"is_collects_selected"`
		ChapterList        interface{} `json:"chapter_list"`
		FeedCommentConfig  struct {
			InputConfigText   string `json:"input_config_text"`
			AuthorAuditStatus int    `json:"author_audit_status"`
			CommonFlags       string `json:"common_flags"`
		} `json:"feed_comment_config"`
		IsImageBeat          bool        `json:"is_image_beat"`
		DislikeDimensionList interface{} `json:"dislike_dimension_list"`
		StandardBarInfoList  interface{} `json:"standard_bar_info_list"`
		PhotoSearchEntrance  struct {
			EcomType int `json:"ecom_type"`
		} `json:"photo_search_entrance"`
		DanmakuControl struct {
			EnableDanmaku      bool   `json:"enable_danmaku"`
			PostPrivilegeLevel int    `json:"post_privilege_level"`
			IsPostDenied       bool   `json:"is_post_denied"`
			PostDeniedReason   string `json:"post_denied_reason"`
			SkipDanmaku        bool   `json:"skip_danmaku"`
			DanmakuCnt         int    `json:"danmaku_cnt"`
			Activities         []struct {
				Id   int `json:"id"`
				Type int `json:"type"`
			} `json:"activities"`
			PassThroughParams string `json:"pass_through_params"`
			SmartModeDecision int    `json:"smart_mode_decision"`
		} `json:"danmaku_control,omitempty"`
		IsLifeItem           bool        `json:"is_life_item"`
		ImageList            interface{} `json:"image_list"`
		ComponentInfoV2      string      `json:"component_info_v2"`
		ItemWarnNotification struct {
			Type    int    `json:"type"`
			Show    bool   `json:"show"`
			Content string `json:"content"`
		} `json:"item_warn_notification"`
		OriginTextExtra      interface{} `json:"origin_text_extra"`
		DisableRelationBar   int         `json:"disable_relation_bar"`
		PackedClips          interface{} `json:"packed_clips"`
		AuthorMaskTag        int         `json:"author_mask_tag"`
		UserRecommendStatus  int         `json:"user_recommend_status"`
		CollectionCornerMark int         `json:"collection_corner_mark"`
		IsSharePost          bool        `json:"is_share_post"`
		ImageComment         struct {
		} `json:"image_comment"`
		VisualSearchInfo struct {
			IsShowImgEntrance  bool `json:"is_show_img_entrance"`
			IsEcomImg          bool `json:"is_ecom_img"`
			IsHighAccuracyEcom bool `json:"is_high_accuracy_ecom"`
			IsHighRecallEcom   bool `json:"is_high_recall_ecom"`
		} `json:"visual_search_info"`
		TtsIdList                  interface{} `json:"tts_id_list"`
		RefTtsIdList               interface{} `json:"ref_tts_id_list"`
		VoiceModifyIdList          interface{} `json:"voice_modify_id_list"`
		RefVoiceModifyIdList       interface{} `json:"ref_voice_modify_id_list"`
		AuthenticationToken        string      `json:"authentication_token"`
		VideoGameDataChannelConfig struct {
		} `json:"video_game_data_channel_config"`
		DislikeDimensionListV2 interface{} `json:"dislike_dimension_list_v2"`
		DistributeCircle       struct {
			DistributeType         int  `json:"distribute_type"`
			CampusBlockInteraction bool `json:"campus_block_interaction"`
			IsCampus               bool `json:"is_campus"`
		} `json:"distribute_circle"`
		ImageCropCtrl    int         `json:"image_crop_ctrl"`
		YummeRecreason   interface{} `json:"yumme_recreason"`
		SlidesMusicBeats interface{} `json:"slides_music_beats"`
		JumpTabInfoList  interface{} `json:"jump_tab_info_list"`
		MediaType        int         `json:"media_type"`
		PlayProgress     struct {
			PlayProgress     int `json:"play_progress"`
			LastModifiedTime int `json:"last_modified_time"`
		} `json:"play_progress"`
		ReplySmartEmojis         interface{} `json:"reply_smart_emojis"`
		ActivityVideoType        int         `json:"activity_video_type"`
		BoostStatus              int         `json:"boost_status"`
		CreateScaleType          interface{} `json:"create_scale_type"`
		EntertainmentProductInfo struct {
			SubTitle   interface{} `json:"sub_title"`
			MarketInfo struct {
				LimitFree struct {
					InFree bool `json:"in_free"`
				} `json:"limit_free"`
				MarketingTag interface{} `json:"marketing_tag"`
			} `json:"market_info"`
		} `json:"entertainment_product_info"`
		Caption       string `json:"caption"`
		ItemTitle     string `json:"item_title"`
		IsUseMusic    bool   `json:"is_use_music"`
		Original      int    `json:"original"`
		XiguaBaseInfo struct {
			Status           int `json:"status"`
			StarAltarOrderId int `json:"star_altar_order_id"`
			StarAltarType    int `json:"star_altar_type"`
			ItemId           int `json:"item_id"`
		} `json:"xigua_base_info"`
		MarkLargelyFollowing bool `json:"mark_largely_following"`
		FriendRecommendInfo  struct {
			FriendRecommendSource            int         `json:"friend_recommend_source"`
			LabelUserList                    interface{} `json:"label_user_list"`
			DisableFriendRecommendGuideLabel bool        `json:"disable_friend_recommend_guide_label"`
		} `json:"friend_recommend_info"`
		GalileoPadTextcrop struct {
			IpadDHCutRatio    []int `json:"ipad_d_h_cut_ratio"`
			IpadDVCutRatio    []int `json:"ipad_d_v_cut_ratio"`
			AndroidDHCutRatio []int `json:"android_d_h_cut_ratio"`
			AndroidDVCutRatio []int `json:"android_d_v_cut_ratio"`
			Version           int   `json:"version"`
		} `json:"galileo_pad_textcrop,omitempty"`
		ImageItemQualityLevel int `json:"image_item_quality_level,omitempty"`
		LifeAnchorShowExtra   struct {
			AnchorType    int    `json:"anchor_type"`
			ShouldShow    bool   `json:"should_show"`
			HasAnchorInfo bool   `json:"has_anchor_info"`
			Extra         string `json:"extra"`
		} `json:"life_anchor_show_extra,omitempty"`
		AnchorInfo struct {
			Type int    `json:"type"`
			Id   string `json:"id"`
			Icon struct {
				Uri     string        `json:"uri"`
				UrlList []interface{} `json:"url_list"`
				Width   int           `json:"width"`
				Height  int           `json:"height"`
				UrlKey  string        `json:"url_key"`
			} `json:"icon"`
			Title     string `json:"title"`
			OpenUrl   string `json:"open_url"`
			WebUrl    string `json:"web_url"`
			MpUrl     string `json:"mp_url"`
			TitleTag  string `json:"title_tag"`
			Content   string `json:"content"`
			StyleInfo struct {
				DefaultIcon string `json:"default_icon"`
				SceneIcon   string `json:"scene_icon"`
				Extra       string `json:"extra"`
			} `json:"style_info"`
			Extra    string `json:"extra"`
			LogExtra string `json:"log_extra"`
		} `json:"anchor_info,omitempty"`
		OriginalAnchorType int `json:"original_anchor_type,omitempty"`
		IsMultiContent     int `json:"is_multi_content,omitempty"`
	} `json:"aweme_list"`
	TimeList []string `json:"time_list"`
	LogPb    struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	RequestItemCursor  int `json:"request_item_cursor"`
	PostSerial         int `json:"post_serial"`
	ReplaceSeriesCover int `json:"replace_series_cover"`
}
