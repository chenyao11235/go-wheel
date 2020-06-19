package main

import (
    "fmt"
    "github.com/spf13/viper"
    "time"
    "wheel/config/useViper/config"
)

func main() {
    config.Init()

    fmt.Println(viper.GetString("env"))
    fmt.Println(viper.GetString("address"))

    for range time.Tick(time.Second) {
    }
}
