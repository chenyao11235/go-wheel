package utils

import (
	"flag"
	"fmt"
	"os"
)

// 如何获取命令行参数
func InitArgs() {
	var (
		isHelp     bool
		configFile string
	)

	flag.BoolVar(&isHelp, "h", false, "查看帮助")
	flag.StringVar(&configFile, "config", "/config.useJsonFile", "指定配置文件")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "需要指定参数, 详见如下: \n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if isHelp {
		flag.Usage()
	}

	if configFile == "" {
		flag.Usage()
	}
}
