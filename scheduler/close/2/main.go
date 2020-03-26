package main

import (
    "fmt"
    "math/rand"
    "time"
)

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

    randStream := newRandStream(done)

    for i := 0; i < 3; i++ {
        fmt.Printf("%d\n", <-randStream)
    }
    close(done)

    time.Sleep(1 * time.Second)
    fmt.Println("do other something")

}
