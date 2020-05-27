package main

import (
    "fmt"
    "wheel/config/useGo-ini/setting"
)

// 通过setting模块来引用各种配置项
func main() {
    fmt.Println(setting.HttpPort)
    fmt.Println(setting.RunMode)
}
