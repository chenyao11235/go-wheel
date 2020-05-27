package main

import (
    "log"
    "wheel/rabbitmq-demo/worker"
)

func main() {
    var (
        err error
    )

    if err = worker.InitRabbitMQ("wheel"); err != nil {
        log.Fatalf("mq初始化错误%s\n", err)
        return
    }

    defer worker.G_rabbitmq.Close()

    worker.G_rabbitmq.Receive(2)
}
