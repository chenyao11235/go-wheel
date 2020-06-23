package main

import (
	"fmt"
	"time"
	"wheel/config/useViper/config"

	"github.com/spf13/viper"
)

// 你这是开什么玩笑
func main() {
	config.Init()

	fmt.Println(viper.GetString("env"))
	fmt.Println(viper.GetString("address"))

	for range time.Tick(time.Second) {
	}
}
