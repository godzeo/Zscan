package main

import (
	logger "Zscan/core/logger"
	options "Zscan/core/option"
	"Zscan/core/runner"
	"fmt"
)

func main() {
	options.ShowBanner()

	// 解析命令行设置
	options := options.ParseOptions()

	// 创建扫描器
	if options.Method == "scan" {

		noahRunner, _ := runner.New(options)
		logger.Infof("初始化runner")

		//noahRunner.RunEnumeration()
		fmt.Println(noahRunner.Options)
	}
}
