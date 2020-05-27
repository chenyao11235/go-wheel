package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
)

const (
	//队列名称
	mqurl = "amqp://guest:guest@10.211.55.5:5672"
)

type RabbitMQ struct {
	connect  *amqp.Connection
	channel  *amqp.Channel
	exchange string
	routeKey string
}

func (rabbitmq *RabbitMQ) Close() {
	rabbitmq.channel.Close()
	rabbitmq.connect.Close()
}

// 订阅模式的生产者
func (rabbitmq *RabbitMQ) Publish(message amqp.Publishing) (err error) {
	err = rabbitmq.channel.ExchangeDeclare(rabbitmq.exchange,
		"topic", // 指定交换机类型，fanout不处理路由键
		true,    // 是否持久化，不持久化，重启之后exchange会消失
		false,   // 如果所有的binding都消失的话，此exchange会被删除
		false,   // 是否是内部专用exchange，是的话，我们不能往里面发送消息
		false,   //
		nil)

	err = rabbitmq.channel.Publish(
		rabbitmq.exchange,
		rabbitmq.routeKey,
		false,
		false,
		message)
	fmt.Printf("发布者发布消息%s\n", string(message.Body))
	return
}

// 订阅模式的消费者
func (rabbitmq *RabbitMQ) Subscribe(n int) (err error) {
	var (
		queue       amqp.Queue
		messageChan <-chan amqp.Delivery
		message     amqp.Delivery
	)

	err = rabbitmq.channel.ExchangeDeclare(
		rabbitmq.exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil)

	queue, err = rabbitmq.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)

	err = rabbitmq.channel.QueueBind(
		queue.Name,
		rabbitmq.routeKey,
		rabbitmq.exchange,
		true,
		nil)

	messageChan, err = rabbitmq.channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	for message = range messageChan {
		fmt.Printf("消费者%d收到消息%s\n", n, string(message.Body))
	}

	return
}

// 构造mq消息
func BuildRabbitMQMessage(content string) (message amqp.Publishing) {
	return amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(content),
	}

}

func InitRabbitMQ(exchange string, routeKey string) (rabbitmq *RabbitMQ, err error) {
	var (
		connect *amqp.Connection
		channel *amqp.Channel
	)

	if connect, err = amqp.Dial(mqurl); err != nil {
		return
	}

	if channel, err = connect.Channel(); err != nil {
		return
	}

	rabbitmq = &RabbitMQ{
		connect:  connect,
		channel:  channel,
		exchange: exchange,
		routeKey: routeKey,
	}

	return
}
