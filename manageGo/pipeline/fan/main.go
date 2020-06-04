package main

import (
    "fmt"
    "math"
    "math/rand"
    "runtime"
    "sync"
    "time"
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

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan interface{} {
    intStream := make(chan interface{})

    go func() {
        defer close(intStream)

        for v := range valueStream {
            select {
            case <-done:
                return
            case intStream <- v.(int):
            }
        }
    }()
    return intStream
}

func primeFinder(done <-chan interface{}, intStream <-chan interface{}) <-chan interface{} {
    primeStream := make(chan interface{})

    go func() {
        defer close(primeStream)
        for {
            select {
            case <-done:
                return
            case item := <-intStream:
                if IfPrime(item.(int)) {
                    primeStream <- item
                }
            }
        }
    }()
    return primeStream
}

func IfPrime(value int) bool {
    if value < 2 {
        return false
    }
    end := int(math.Sqrt(float64(value)))
    for i := 2; i <= end; i++ {
        if value%i == 0 {
            return false
        }
    }
    return true
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
    var wg sync.WaitGroup
    multiplexedStream := make(chan interface{})

    multiplex := func(c <-chan interface{}) {
        defer wg.Done()
        for i := range c {
            select {
            case <-done:
                return
            case multiplexedStream <- i:
            }
        }
    }

    wg.Add(len(channels))
    for _, c := range channels {
        go multiplex(c)
    }

    go func() {
        wg.Wait()
        close(multiplexedStream)
    }()

    return multiplexedStream
}

func main() {
    done := make(chan interface{})
    defer close(done)

    rand := func() interface{} { return rand.Intn(50000000) }

    start := time.Now()
    randIntStream := toInt(done, repeatFn(done, rand))

    numsFinders := runtime.NumCPU()
    finders := make([]<-chan interface{}, numsFinders)

    fmt.Println("Primes:")

    for i := 0; i < numsFinders; i++ {
        finders[i] = primeFinder(done, randIntStream)
    }

    //for prime := range take(done, primeFinder(done, randIntStream), 1000000) {
    //    fmt.Printf("\t%d\n", prime)
    //}

    for prime := range take(done, fanIn(done, finders...), 1000000) {
        fmt.Printf("\t%d\n", prime)
    }

    fmt.Printf("Search took: %v\n", time.Since(start))
}
