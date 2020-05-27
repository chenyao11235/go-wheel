package main

import (
    "fmt"
    "math/rand"
)

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
    valueStream := make(chan interface{})
    go func() {
        defer close(valueStream)
        for {
            for _, v := range values {
                select {
                case <-done:
                    return
                case valueStream <- v:
                }
            }
        }
    }()
    return valueStream
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
    valueStream := make(chan interface{})
    go func() {
        defer close(valueStream)
        for {
            select {
            case <-done:
                return
            case valueStream <- fn():
            }
        }
    }()
    return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
    takeStream := make(chan interface{})

    go func() {
        defer close(takeStream)
        for i := 0; i < num; i++ {
            select {
            case <-done:
                return
            case item := <-valueStream:
                takeStream <- item
            }
        }
    }()
    return takeStream
}

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
    stringStream := make(chan string)

    go func() {
        defer close(stringStream)

        for v := range valueStream {
            select {
            case <-done:
                return
            case stringStream <- v.(string):
            }
        }
    }()
    return stringStream

}

func main() {
    done := make(chan interface{})
    defer close(done)

    rand := func() interface{} { return rand.Intn(50000000) }

    for item := range take(done, repeat(done, 1), 10) {
        fmt.Println(item)
    }
    for item := range take(done, repeatFn(done, rand), 10) {
        fmt.Println(item)
    }
    for item := range toString(done, take(done, repeat(done, "I", "am"), 5)) {
        fmt.Println(item)
    }
}
