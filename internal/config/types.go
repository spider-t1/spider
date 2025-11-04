package config

type Config struct {
	System  System  `json:"system"`
	Logging Logging `json:"logging"`
	Cookie  Cookie  `json:"cookie"`
}

type Cookie struct {
	Douyin string `json:"douyin"`
}

type System struct {
	Name      string `json:"name"`
	Env       string `json:"env"`
	Port      string `json:"port"`
	Migration bool   `json:"migration"`
	Chrome    string `json:"chrome"`
}

// Logging 日志配置
type Logging struct {
	Level        string        `json:"level"`         // 日志级别: debug, info, warn, error
	Format       string        `json:"format"`        // 日志格式: json, text
	MaxAge       string        `json:"max_age"`       // 最大保留时间，如"30d"
	RotationTime string        `json:"rotation_time"` // 轮转时间，如"24h"
	Output       []LogOutput   `json:"output"`        // 输出配置
	Categories   LogCategories `json:"categories"`    // 分类日志配置
}

// LogOutput 日志输出配置
type LogOutput struct {
	Type    string `json:"type"`    // 输出类型: file, console
	Path    string `json:"path"`    // 文件路径（仅file类型）
	Colored bool   `json:"colored"` // 是否彩色输出（仅console类型）
}

// LogCategories 分类日志配置
type LogCategories struct {
	Access LogCategory `json:"access"` // 访问日志
	Slow   LogCategory `json:"slow"`   // 慢日志
	Error  LogCategory `json:"error"`  // 错误日志
}

// LogCategory 单个分类日志配置
type LogCategory struct {
	Enabled   bool   `json:"enabled"`   // 是否启用
	Path      string `json:"path"`      // 日志路径
	Threshold int    `json:"threshold"` // 阈值（仅slow类型使用）
}
