package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"wheel/rabbitmq-demo/worker"
	"time"
)

func main() {
	var (
		err     error
		message amqp.Publishing
	)

	if err = worker.InitRabbitMQ("wheel"); err != nil {
		log.Fatalf("mq初始化错误%s\n", err)
		return
	}

	defer worker.G_rabbitmq.Close()

	for i := 1; i <= 10; i++ {
		message = worker.BuildRabbitMQMessage(fmt.Sprintf("消息%d", i))
		_ = worker.G_rabbitmq.Push(message)
		time.Sleep(2 * time.Second)
	}
}
