package main

import (
    "fmt"
    "time"
)

/*
   通过waitgroup和 带缓冲channle的结合实现了一种类似信号量的机制
   不仅可以实现main groutine和child goroutine的同步也可以实现并发数量的控制
*/

var sema = NewSemaphore(3)

func Read(i int) {
    defer sema.Done()
    sema.Add(1)

    fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
    time.Sleep(time.Second)
}

func main() {
    userCount := 10
    for i := 0; i < userCount; i++ {
        go Read(i)
    }

    sema.Wait()
}
