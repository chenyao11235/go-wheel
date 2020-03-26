package main

import (
	rabbitMQ "rabbitMQ/routing"
)

func main() {
	var (
		err      error
		rabbitmq *rabbitMQ.RabbitMQ
	)

	// 只有绑定了同一个交换机和路由键和队列才会收到消息，而订阅模式只是绑定同一个交换机就可以
	if rabbitmq, err = rabbitMQ.InitRabbitMQ("wheel1", "one"); err != nil {
		return
	}

	//defer rabbitmq.Close()

	_ = rabbitmq.Subscribe(1)
}
