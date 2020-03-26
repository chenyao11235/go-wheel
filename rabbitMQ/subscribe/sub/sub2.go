package main

import (
	"log"
	"rabbitMQ/subscribe"
)

func main() {
	var (
		err error
	)

	if err = rabbitMQ.InitRabbitMQ("wheel"); err != nil {
		log.Fatalf("初始化mq失败 %s", err)
		return
	}

	defer rabbitMQ.G_rabbitmq.Close()

	_ = rabbitMQ.G_rabbitmq.Subscribe(1)
}
