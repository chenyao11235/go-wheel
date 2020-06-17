package main

import (
    "math/rand"
    "strconv"
    "testing"
    "time"
)

func TestAFunc(t *testing.T) {
    scheme := strconv.FormatInt(time.Now().UnixNano(), 36)
    t.Log(scheme)
}

type Monitor struct {
    timeoutChane chan int64
}

func TestAfter(t *testing.T) {
    m := &Monitor{timeoutChane: make(chan int64)}
    for i := 0; i < 10; i++ {
        n := rand.Int63n(600000)
        t.Log(n)
        go func(n int64) {
            select {
            case <-time.After(time.Second * 3):
                m.timeoutChane <- n
            }
        }(n)

    }

    t.Log("start monitor...")

    for {
        select {
        case id := <-m.timeoutChane:
            t.Logf("%d timeout", id)
        default:
        }
    }
}
