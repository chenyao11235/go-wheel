package main

// 改包展示了如何载入json形式的配置文件
import (
    "encoding/json"
    "fmt"
    "os"
)

type Configuration struct {
    Address      string
    ReadTimeout  int64
    WriteTimeout int64
    Static       string
}

// 单例模式，全局变量
var (
    G_config *Configuration
)

func InitConfig(path string) (err error) {
    var (
        file    *os.File
        decoder *json.Decoder
    )

    if file, err = os.Open(path); err != nil {
        return
    }

    decoder = json.NewDecoder(file)
    if err = decoder.Decode(G_config); err != nil {
        return
    }

    return
}

func main() {
    var (
        err error
    )
    // 载入配置文件
    if err = InitConfig("config.json"); err != nil {
        goto ERR
    }

ERR:
    fmt.Println(err)
    return
}
