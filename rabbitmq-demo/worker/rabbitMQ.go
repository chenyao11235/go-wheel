package worker

import (
    "fmt"
    "github.com/streadway/amqp"
)

const (
    mqUrl = "amqp://guest:guest@10.211.55.5:5672"
)

type RabbitMQ struct {
    connect *amqp.Connection
    channel *amqp.Channel
    queue   amqp.Queue
}

var (
    G_rabbitmq *RabbitMQ
)

func (rabbitmq *RabbitMQ) Close() {
    rabbitmq.channel.Close()
    rabbitmq.connect.Close()
}

// simple模式的生产者
func (rabbitmq *RabbitMQ) Push(message amqp.Publishing) (err error) {
    err = rabbitmq.channel.Publish(
        "",
        rabbitmq.queue.Name,
        false, // 如果不能能找到合适的目标queue,则通过Channel.NotifyReturn将消息返回给发布者
        false, message) // 如果消费者不能立即消费这个消息，则通过Channel.NotifyReturn将消息返回给发布者
    return
}

// simple模式的消费者
func (rabbitmq *RabbitMQ) Receive(n int) {
    var (
        messgeChan <-chan amqp.Delivery
        message    amqp.Delivery
    )

    messgeChan, _ = rabbitmq.channel.Consume(
        rabbitmq.queue.Name,
        "",    // 消费者标签
        true,  // 自动应答，消费者收到消息之后通知生产者者，生产者将此消息删除
        false, // 排他消费者,即这个队列只能由一个消费者消费.适用于任务不允许进行并发处理的情况下.比如系统对接
        false, // 这个功能属于AMQP的标准,但是rabbitMQ并没有做实现.
        false, // 不返回执行结果,但是如果排他开启的话,则必须需要等待结果的,如果两个一起开就会报错
        nil)

    for message = range messgeChan {
        fmt.Printf("消费者%d收消息 %s\n", n, string(message.Body))
    }
}

// 构造mq消息
func BuildRabbitMQMessage(content string) (message amqp.Publishing) {
    return amqp.Publishing{
        ContentType: "text/plain",
        Body:        []byte(content),
    }

}

func InitRabbitMQ(queueName string) (err error) {
    var (
        connect *amqp.Connection
        channel *amqp.Channel
        queue   amqp.Queue
    )

    if connect, err = amqp.Dial(mqUrl); err != nil {
        return
    }

    if channel, err = connect.Channel(); err != nil {
        return
    }

    if queue, err = channel.QueueDeclare(queueName,
        false,
        false,
        false,
        false,
        nil); err != nil {
        return
    }

    G_rabbitmq = &RabbitMQ{
        connect: connect,
        channel: channel,
        queue:   queue,
    }

    return
}
