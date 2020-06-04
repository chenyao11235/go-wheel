package main

import (
    "fmt"
    "time"
)

func main() {
    pool, err := NewPool(10)
    if err != nil {
        panic(err)
    }

    for i := 0; i < 20; i++ {
        err := pool.Put(&Task{
            Handler: func(params ...interface{}) {
                fmt.Println(params[0])
            },
        })
        if err != nil {
            fmt.Println(err)
        }
    }
    time.Sleep(10 * time.Second)
}
