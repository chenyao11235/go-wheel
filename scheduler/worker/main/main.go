package main

import (
    "log"
    "sync"
    "time"
    "wheel/scheduler/worker"
)

var names = []string{
    "steve",
    "bob",
    "mary",
    "therese",
    "jason",
}

type namePrinter struct {
    name string
}

func (m *namePrinter) Task() {
    log.Println(m.name)
    time.Sleep(time.Second)
}

func main() {
    var (
        wg sync.WaitGroup
    )
    p := worker.New(2)

    wg.Add(100 * len(names))

    for i := 0; i < 100; i++ {
        for _, name := range names {
            np := namePrinter{name: name}
            go func() {
                p.Run(&np)
                wg.Done()
            }()
        }
    }

    wg.Wait()
    p.Shutdown()
}
