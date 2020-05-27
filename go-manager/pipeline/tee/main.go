package main

import "fmt"

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
func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
    valueStream := make(chan interface{})
    go func() {
        defer close(valueStream)

        for {
            select {
            case <-done:
                return
            case v, ok := <-c:
                if !ok {
                    return
                }
                select {
                case valueStream <- v:
                case <-done:
                }
            }
        }
    }()
    return valueStream
}

func tee(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
    out1 := make(chan interface{})
    out2 := make(chan interface{})

    go func() {
        defer close(out1)
        defer close(out2)

        for val := range orDone(done, in) {
            var out1, out2 = out1, out2
            for i := 0; i < 2; i++ {
                select {
                case <-done:
                case out1 <- val:
                    out1 = nil
                case out2 <- val:
                    out2 = nil
                }
            }
        }
    }()
    return out1, out2
}

func main() {
    done := make(chan interface{})
    defer close(done)

    out1, out2 := tee(done, take(done, repeat(done, 1, 2), 10))

    for val1 := range out1 {
        fmt.Printf(":out:%v, out2:%v\n", val1, <-out2)
    }

}
