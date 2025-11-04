package types_douyin

import "spider/internal/app/types/types_common"

// CommentListReq 抖音评论列表请求参数
type CommentListReq struct {
	types_common.BaseListParam
	AwemeId string `form:"awemeId" binding:"required"`
	Cursor  string `form:"cursor"`
	Count   string `form:"count"`
}

func (r *CommentListReq) Adjust() {
	if r.Cursor == "" {
		r.Cursor = "0"
	}
	if r.Count == "" {
		r.Count = "10"
	}
}

// CommentReplyListReq 抖音评论回复列表请求参数
type CommentReplyListReq struct {
	types_common.BaseListParam
	ItemId    string `form:"item_id" binding:"required"`
	CommentId string `form:"comment_id" binding:"required"`
	Cursor    string `form:"cursor"`
	Count     string `form:"count"`
}

func (r *CommentReplyListReq) Adjust() {
	if r.Cursor == "" {
		r.Cursor = "0"
	}
	if r.Count == "" {
		r.Count = "10"
	}
}

type CommentData struct {
	StatusCode               int                  `json:"status_code"`
	Comments                 []Comment            `json:"comments"`
	Cursor                   int                  `json:"cursor"`
	HasMore                  int                  `json:"has_more"`
	ReplyStyle               int                  `json:"reply_style"`
	Total                    int                  `json:"total"` // 总评论数
	Extra                    Extra                `json:"extra"`
	LogPb                    LogPb                `json:"log_pb"`
	HotsoonFilteredCount     int                  `json:"hotsoon_filtered_count"`
	UserCommented            int                  `json:"user_commented"`
	FastResponseComment      FastResponseComment  `json:"fast_response_comment"`
	CommentConfig            CommentConfig        `json:"comment_config"`
	GeneralCommentConfig     GeneralCommentConfig `json:"general_comment_config"`
	ShowManagementEntryPoint int                  `json:"show_management_entry_point"`
	CommentCommonData        string               `json:"comment_common_data"`
	FoldedCommentCount       int                  `json:"folded_comment_count"`
}

type Comment struct {
	Cid               string         `json:"cid"`
	Text              string         `json:"text"`
	AwemeId           string         `json:"aweme_id"`
	CreateTime        int            `json:"create_time"`
	DiggCount         int            `json:"digg_count"`
	Status            int            `json:"status"`
	User              User           `json:"user"`
	ReplyId           string         `json:"reply_id"`
	UserDigged        int            `json:"user_digged"`
	ReplyComment      []ReplyComment `json:"reply_comment"`
	TextExtra         []interface{}  `json:"text_extra"`
	LabelText         string         `json:"label_text"`
	LabelType         int            `json:"label_type"`
	ReplyCommentTotal int            `json:"reply_comment_total"`
	ReplyToReplyId    string         `json:"reply_to_reply_id"`
	IsAuthorDigged    bool           `json:"is_author_digged"`
	StickPosition     int            `json:"stick_position"`
	UserBuried        bool           `json:"user_buried"`
	LabelList         []Label        `json:"label_list"`
	IsHot             bool           `json:"is_hot"`
	TextMusicInfo     interface{}    `json:"text_music_info"`
	ImageList         interface{}    `json:"image_list"`
	IsNoteComment     int            `json:"is_note_comment"`
	IpLabel           string         `json:"ip_label"`
	CanShare          bool           `json:"can_share"`
	ItemCommentTotal  int            `json:"item_comment_total"`
	Level             int            `json:"level"`
	VideoList         interface{}    `json:"video_list"`
	SortTags          string         `json:"sort_tags"`
	IsUserTendToReply bool           `json:"is_user_tend_to_reply"`
	ContentType       int            `json:"content_type"`
	IsFolded          bool           `json:"is_folded"`
	EnterFrom         string         `json:"enter_from"`
	Sticker           *Sticker       `json:"sticker,omitempty"`
}

type ReplyComment = Comment // 如结构一致，可复用 Comment

type User struct {
	Uid                      string        `json:"uid"`
	Nickname                 string        `json:"nickname"`
	SecUid                   string        `json:"sec_uid"`
	ShortId                  string        `json:"short_id"`
	AvatarThumb              AvatarThumb   `json:"avatar_thumb"`
	CustomVerify             string        `json:"custom_verify"`
	IsAdFake                 bool          `json:"is_ad_fake"`
	Region                   string        `json:"region"`
	Status                   int           `json:"status"`
	CommentSetting           int           `json:"comment_setting"`
	IsStar                   bool          `json:"is_star"`
	CloseFriendType          int           `json:"close_friend_type"`
	DisableImageCommentSaved int           `json:"disable_image_comment_saved"`
	BanUserFunctions         []interface{} `json:"ban_user_functions"`
	AwemeControl             AwemeControl  `json:"aweme_control"`
	EnterpriseVerifyReason   string        `json:"enterprise_verify_reason"`
	CommerceUserLevel        int           `json:"commerce_user_level"`
	// 其他字段省略，保留 interface{} 类型字段不建议列全
}

type AvatarThumb struct {
	Uri     string   `json:"uri"`
	UrlList []string `json:"url_list"`
	Width   int      `json:"width"`
	Height  int      `json:"height"`
}

type AwemeControl struct {
	CanForward     bool `json:"can_forward"`
	CanShare       bool `json:"can_share"`
	CanComment     bool `json:"can_comment"`
	CanShowComment bool `json:"can_show_comment"`
}

type Label struct {
	Type int    `json:"type"`
	Text string `json:"text"`
}

type Sticker struct {
	Id              int64  `json:"id"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	StaticUrl       Media  `json:"static_url"`
	AnimateUrl      Media  `json:"animate_url"`
	StickerType     int    `json:"sticker_type"`
	OriginPackageId int64  `json:"origin_package_id"`
	IdStr           string `json:"id_str"`
	AuthorSecUid    string `json:"author_sec_uid"`
}

type Media struct {
	Uri     string   `json:"uri"`
	UrlList []string `json:"url_list"`
	Width   int      `json:"width"`
	Height  int      `json:"height"`
}

type Extra struct {
	Now          int64       `json:"now"`
	FatalItemIds interface{} `json:"fatal_item_ids"`
}

type LogPb struct {
	ImprId string `json:"impr_id"`
}

type FastResponseComment struct {
	ConstantResponseWords []string `json:"constant_response_words"`
	TimedResponseWords    []string `json:"timed_response_words"`
}

type CommentConfig struct{}
type GeneralCommentConfig struct{}
