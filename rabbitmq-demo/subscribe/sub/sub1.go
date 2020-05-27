package main

import (
    "log"
    "wheel/rabbitmq-demo/subscribe"
)

const (
    exchangeName = "wheel"
)

func main() {
    var (
        err error
    )

    if err = rabbitMQ.InitRabbitMQ(exchangeName); err != nil {
        log.Fatalf("初始化mq失败 %s", err)
        return
    }

    defer rabbitMQ.G_rabbitmq.Close()

    _ = rabbitMQ.G_rabbitmq.Subscribe(1)
}
