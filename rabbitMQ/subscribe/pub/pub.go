package main

import (
	"log"
	"rabbitMQ/subscribe"
	"strconv"
	"time"
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

	for i := 1; i < 10; i++ {
		message := rabbitMQ.BuildRabbitMQMessage("消息" + strconv.Itoa(i))
		_ = rabbitMQ.G_rabbitmq.Publish(message)
		time.Sleep(1 * time.Second)
	}
}
