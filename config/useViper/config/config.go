package config

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/joho/godotenv"
    "github.com/pkg/errors"
    "github.com/spf13/viper"
    "log"
)

/* viper 读取配置信息的优先级
    1  viper set
    2  命令行
    3 配置文件
    4 环境变量
*/



func Init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }
    // viper加载环境变量
    viper.AutomaticEnv()
    // 设置配置文件的名字和类型
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    // 可以添加多个配置文件的路径
    viper.AddConfigPath("./conf")
    viper.AddConfigPath(".")
    // 读取配置文件
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            log.Fatal(errors.Wrap(err, "config file not found"))
        } else {
            log.Fatal(err)
        }
    }
    //
    viper.WatchConfig()
    viper.OnConfigChange(func(in fsnotify.Event) {
        fmt.Println("config file changed")
    })
}
