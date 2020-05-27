package main

import (
    "strconv"
    "time"
    "wheel/rabbitmq-demo/routing"
)

func main() {
    var (
        err         error
        rabbitmqOne *rabbitMQ.RabbitMQ
        rabbitmqTwo *rabbitMQ.RabbitMQ
    )

    // 只有绑定了同一个交换机和路由键和队列才会收到消息，而订阅模式只是绑定同一个交换机就可以
    if rabbitmqOne, err = rabbitMQ.InitRabbitMQ("wheel1", "one"); err != nil {
        return
    }

    if rabbitmqTwo, err = rabbitMQ.InitRabbitMQ("wheel1", "two"); err != nil {
        return
    }

    defer rabbitmqOne.Close()
    defer rabbitmqTwo.Close()

    for i := 1; i < 10; i++ {
        message1 := rabbitMQ.BuildRabbitMQMessage("one 消息" + strconv.Itoa(i))
        message2 := rabbitMQ.BuildRabbitMQMessage("two 消息" + strconv.Itoa(i))
        _ = rabbitmqOne.Publish(message1)
        _ = rabbitmqTwo.Publish(message2)
        time.Sleep(1 * time.Second)
    }
}
