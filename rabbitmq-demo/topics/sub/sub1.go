package main

import (
    rabbitMQ "wheel/rabbitmq-demo/topics"
)

func main() {
    var (
        err      error
        rabbitmq *rabbitMQ.RabbitMQ
    )

    // "#"表示通配所有的队列
    if rabbitmq, err = rabbitMQ.InitRabbitMQ("wheel2", "#"); err != nil {
        return
    }

    defer rabbitmq.Close()

    _ = rabbitmq.Subscribe(1)
}
