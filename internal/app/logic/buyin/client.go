package buyin

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/dop251/goja"
	"io"
	"net/http"
	"net/url"
	"os"
	"spider/pkg/utils"
	"strings"
)

// ClientConfig 客户端配置结构体
type ClientConfig struct {
	Cookie    string // 完整的cookie字符串
	EWID      string `json:"ewid,omitempty"` // ewid参数
	VerifyFp  string // verifyFp参数
	Fp        string // fp参数
	MsToken   string // msToken参数
	UserAgent string // User-Agent，可选，有默认值
}

// JXClient 抖音精选联盟API客户端
type JXClient struct {
	config     ClientConfig
	httpClient *http.Client
}

// NewJXClient 创建新的客户端实例
func NewJXClient(config ClientConfig) *JXClient {
	// 设置默认User-Agent
	if config.UserAgent == "" {
		config.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0"
	}

	return &JXClient{
		config:     config,
		httpClient: &http.Client{},
	}
}

// generateABogus 生成a_bogus参数
func (c *JXClient) generateABogus(pa, da string) (string, error) {
	jsContent, err := os.ReadFile("./longab_20.js")
	if err != nil {
		return "", fmt.Errorf("读取JavaScript文件失败: %v", err)
	}

	vm := goja.New()

	console := vm.NewObject()
	console.Set("log", func(args ...interface{}) { fmt.Println(args...) })
	console.Set("error", func(args ...interface{}) { fmt.Println("ERROR:", args) })
	console.Set("warn", func(args ...interface{}) { fmt.Println("WARN:", args) })
	console.Set("info", func(args ...interface{}) { fmt.Println("INFO:", args) })
	vm.Set("console", console)

	jsCode := string(jsContent)
	if strings.Contains(jsCode, "console.log(enc('','',''))") {
		jsCode = strings.Replace(jsCode, "console.log(enc('','',''))", "// console.log(enc('','',''))", 1)
	}

	_, err = vm.RunString(jsCode)
	if err != nil {
		return "", fmt.Errorf("执行JavaScript代码失败: %v", err)
	}

	var enc func(string, string, string) string
	if err = vm.ExportTo(vm.Get("enc"), &enc); err != nil {
		return "", fmt.Errorf("导出enc函数失败: %v", err)
	}

	return enc(pa, da, c.config.UserAgent), nil
}

// doRequest 执行HTTP请求的通用方法
func (c *JXClient) doRequest(method, baseURL string, headers map[string]string, queryParams map[string]string, body []byte) (string, string, error) {
	// 构建查询参数字符串
	var paramPairs []string
	for key, value := range queryParams {
		paramPairs = append(paramPairs, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
	}
	pa := strings.Join(paramPairs, "&")

	// 生成a_bogus
	da := ""
	if body != nil {
		da = string(body)
	}
	fmt.Println("pa:", pa)
	fmt.Println("da:", da)
	aBogus, err := c.generateABogusV2(pa, da)
	if err != nil {
		return "", "", fmt.Errorf("生成a_bogus失败: %v", err)
	}

	// 构建完整URL
	fullURL := fmt.Sprintf("%s?%s&a_bogus=%s", baseURL, pa, url.QueryEscape(aBogus))

	// 创建请求
	var req *http.Request
	if body != nil {
		req, err = http.NewRequest(method, fullURL, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, fullURL, nil)
	}
	if err != nil {
		return "", "", fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Cookie", c.config.Cookie)

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应，处理可能的gzip压缩
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return "", "", fmt.Errorf("创建gzip读取器失败: %v", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	responseBody, err := io.ReadAll(reader)
	if err != nil {
		return "", "", fmt.Errorf("读取响应失败: %v", err)
	}

	return string(responseBody), resp.Status, nil
}

// generateABogusV2 生成a_bogus参数 - V2版本，使用Go实现
func (c *JXClient) generateABogusV2(pa, da string) (string, error) {
	// 直接调用Go版本的Enc函数
	result := utils.Enc(pa, da, c.config.UserAgent)
	return result, nil
}

func GenerateABogusV2(pa, da, uAgent string) (string, error) {
	// 直接调用Go版本的Enc函数
	result := utils.Enc(pa, da, uAgent)
	return result, nil
}
