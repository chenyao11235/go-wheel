package main

import (
    "fmt"
    "math/rand"
    "time"
)

func doWork(done <-chan interface{}, strings <-chan string) <-chan interface{} {
    terminated := make(chan interface{})
    go func() {
        defer fmt.Println("doWork exit")
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

func newRandStream(done <-chan interface{}) <-chan int {
    randStream := make(chan int)

    go func() {
        defer fmt.Println("newRandStream closure exited")
        defer close(randStream)

        for {
            select {
            case <-done:
                return
            case randStream <- rand.Int():
            }
        }
    }()

    return randStream
}

func main() {
    done := make(chan interface{})
    // nil的chan会造成case的永久阻塞，但是超时机制让阻塞的goroutine超时取消
    terminated := doWork(done, nil)

    go func() {
        time.Sleep(3 * time.Second)
        fmt.Println("Canceling doWork goroutine")
        close(done)
    }()

    fmt.Println(<-terminated)
    fmt.Println("Done")

}
