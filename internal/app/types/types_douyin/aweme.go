package types_douyin

import "spider/internal/app/types/types_common"

// AwemeDetailReq 抖音作品详情请求参数
type AwemeDetailReq struct {
    types_common.BaseListParam
    AwemeId string `form:"aweme_id" binding:"required"`
}

func (r *AwemeDetailReq) Adjust() {}