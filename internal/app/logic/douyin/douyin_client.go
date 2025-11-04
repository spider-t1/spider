package douyin

import (
	"context"
	"spider/internal/app/types/types_douyin"
	"spider/internal/config"
)

type DouyinClient struct {
	Cookie    string
	UserAgent string
}

type IDouyinClient interface {
	DouyinAwemeDetail(awemeId string) (string, error)
	DouyinComment(ctx context.Context, req *types_douyin.CommentListReq) (string, error)
	DouyinCommentReply(itemId, commentId, cursor, count string) (string, string, error)
	DouyinUserInfo(secUserId string) (*types_douyin.DouyinUserInfoResp, string, error)
	DouyinUserSearch(keyword string) (*types_douyin.DouyinSearchResp, error)
	DouyinUserVideo(secId string) (string, string, error)
}

func NewDouyinClient(userAgent string) *DouyinClient {
	return &DouyinClient{
		Cookie: func() string {
			return config.Cfg.Cookie.Douyin
		}(),
		UserAgent: func() string {
			if userAgent == "" {
				return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36"
			}
			return userAgent
		}(),
	}
}
