package main

import (
    "fmt"
    "time"
)

// 实现在主goutine中控制child goroutine的运行时间，超时退出
// 思路就是： 父goroutine传给 child goroutine一个channel 通过这个channel实现信号的传递

func doWork(done <-chan interface{}, strings <-chan string) <-chan interface{} {
    terminated := make(chan interface{})
    go func() {
        defer fmt.Println("task timeout...")
        defer close(terminated)

        for {
            select {
            case <-done:
                return
            case s := <-strings:
                fmt.Println(s)
            }
        }
    }()
    return terminated
}

func main() {
    done := make(chan interface{})

    // nil的chan会造成case的永久阻塞，但是超时机制让阻塞的goroutine超时取消
    terminated := doWork(done, nil)

    // 这个goroutine作用是作为 一个 timeout的计时器
    go func() {
        time.Sleep(3 * time.Second)
        close(done)
    }()

    // 这里会阻塞，直到main goroutine收到child goroutine的timeout信号
    <-terminated
    fmt.Println("Done")

}
