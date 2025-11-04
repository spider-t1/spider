package cmd

import (
	"fmt"
	"os/exec"
	"spider/pkg/logger"
)

// SwagInitCmd 生成swagger文档的命令
func SwagInitCmd() {
	// 执行swag init命令
	logger.Logger.Info("正在生成swagger文档...")
	cmd := exec.Command("swag", "init")
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(out)
		return
	}
}
