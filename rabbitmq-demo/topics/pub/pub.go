package main

import (
    "strconv"
    "time"
    "wheel/rabbitmq-demo/topics"
)

func main() {
    var (
        err         error
        rabbitmqOne *rabbitMQ.RabbitMQ
        rabbitmqTwo *rabbitMQ.RabbitMQ
    )

    if rabbitmqOne, err = rabbitMQ.InitRabbitMQ("wheel2", "dispatcher.golang.one"); err != nil {
        return
    }

    if rabbitmqTwo, err = rabbitMQ.InitRabbitMQ("wheel2", "dispatcher.golang.two"); err != nil {
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
