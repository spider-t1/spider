package types_douyin

type DouyinEmojiListResp struct {
    StatusCode int         `json:"status_code"`
    Code       int         `json:"code"`
    Msg        string      `json:"msg"`
    Message    string      `json:"message"`
    Data       interface{} `json:"data"`
    Extra      interface{} `json:"extra"`
}

