package main

import (
    "context"
    "fmt"
    "time"
)

type message struct {
    responseChan chan<- int
    parameter    string
    ctx          context.Context
}

func ProcessMessages(worker <-chan message) {
    for job := range worker {
        select {
        case <-job.ctx.Done():
            continue
        default:
        }

        hardToCalculate := len(job.parameter)
        time.Sleep(5 * time.Second)
        select {
        case <-job.ctx.Done():
        case job.responseChan <- hardToCalculate:
        }
    }
}

func NewRequest(ctx context.Context, input string, q chan<- message) {
    r := make(chan int)
    select {
    case <-ctx.Done():
        fmt.Println("刚开始就结束了...")
    case q <- message{
        responseChan: r,
        parameter:    input,
        ctx:          ctx,
    }:
    }

    select {
    case out := <-r:
        fmt.Printf("The len of %s is %d\n", input, out)
    case <-ctx.Done():
        fmt.Println("我被通知退出了...")
    }
}

func main() {
    q := make(chan message, 2)
    go ProcessMessages(q)
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    go NewRequest(ctx, "hi", q)
    go NewRequest(ctx, "hello", q)
    go func() {
        select {
        case <-ctx.Done():
            fmt.Println("超时时间到了...")
        }
    }()

    ticker := time.NewTicker(1 * time.Second)
    for {
        select {
        case <-ticker.C:
        }
    }
}
