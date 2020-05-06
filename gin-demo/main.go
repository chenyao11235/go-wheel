package main

import (
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/spf13/pflag"
    "github.com/spf13/viper"
    "log"
    "net/http"
    "time"
    "wheel/gin-demo/config"
    "wheel/gin-demo/logger"
    "wheel/gin-demo/router"
)

// 一条典型的curl命令
// curl -v -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/user -d'{"username":"admin","password":"admin1234"}'

var (
    cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
    pflag.Parse()
    if err := config.Init(*cfg); err != nil {
        panic(err)
    }

    logger.InitLog()

    gin.SetMode(viper.GetString("runmode"))
    g := gin.New()

    middlewares := []gin.HandlerFunc{}

    router.Load(
        g,
        middlewares...,
    )

    // Ping the server to make sure the router is working.
    go func() {
        if err := pingServer(); err != nil {
            log.Fatal("The router has no response, or it might took too long to start up.", err)
        }
        log.Print("The router has been deployed successfully.")
    }()

    logger.Log.Infof("Start to listening the incoming requests on http address: %s", ":8080")
    logger.Log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
    for i := 0; i < viper.GetInt("max_ping_count"); i++ {
        // Ping the server by sending a GET request to `/health`.
        resp, err := http.Get(viper.GetString("url") + "/sd/health")
        if err == nil && resp.StatusCode == 200 {
            return nil
        }

        // Sleep for a second to continue the next ping.
        log.Print("Waiting for the router, retry in 1 second.")
        time.Sleep(time.Second)
    }
    return errors.New("Cannot connect to the router.")
}
