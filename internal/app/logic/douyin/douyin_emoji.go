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
)

func (c *DouyinClient) DouyinEmojiList(ctx context.Context, webRid string) (*types_douyin.DouyinEmojiListResp, error) {
    client := &http.Client{}

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

    pa := fmt.Sprintf("aid=%s&app_name=%s&live_id=%s&device_platform=%s&language=%s&enter_from=%s&cookie_enabled=%s&screen_width=%s&screen_height=%s&browser_language=%s&browser_platform=%s&browser_name=%s&browser_version=%s&msToken=%s",
        url.QueryEscape("6383"),
        url.QueryEscape("douyin_web"),
        url.QueryEscape("1"),
        url.QueryEscape("web"),
        url.QueryEscape("zh-CN"),
        url.QueryEscape("web_live"),
        url.QueryEscape("true"),
        url.QueryEscape("1080"),
        url.QueryEscape("1920"),
        url.QueryEscape("zh-CN"),
        url.QueryEscape("Win32"),
        url.QueryEscape("Edge"),
        url.QueryEscape("142.0.0.0"),
        url.QueryEscape("pJn9-fCr6qA-f-K490lrSBXLdwEan5CzHoVN_f-dSPn-Q2zbgTPqmQCDDotElWwBMhUI71MflIYaGmLGZ4sgOZ3NVcP47jbWgoMU6QBrFe5IAnG02UKVDtAz3qH4yFcTiycAxLsGCZnVFH7v0XIrYd9ctfhzLvUwlv0y4p5wjb2g"),
    )

    da := ""

    aBogus, err := c.generateABogusV2(pa, da)
    if err != nil {
        return nil, fmt.Errorf("生成a_bogus失败: %v", err)
    }

    baseURL := "https://live.douyin.com/aweme/v1/web/emoji/list"
    fullURL := fmt.Sprintf("%s?%s&a_bogus=%s", baseURL, pa, url.QueryEscape(aBogus))

    req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
        return nil, fmt.Errorf("创建请求失败: %v", err)
    }

    for k, v := range headers {
        req.Header.Set(k, v)
    }
    req.Header.Set("Cookie", c.Cookie)

    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("请求失败: %v", err)
    }
    defer resp.Body.Close()

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

    var out types_douyin.DouyinEmojiListResp
    if err := json.Unmarshal(body, &out); err != nil {
        return nil, fmt.Errorf("解析响应JSON失败: %v", err)
    }

    return &out, nil
}

