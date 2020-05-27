package main

import "sync"

type Semaphore struct {
    c  chan struct{}
    wg *sync.WaitGroup
}

func NewSemaphore(maxSize int) *Semaphore {
    return &Semaphore{
        c:  make(chan struct{}, maxSize),
        wg: new(sync.WaitGroup),
    }
}

func (s *Semaphore) Add(delta int) {
    s.wg.Add(delta)
    for i := 0; i < delta; i++ {
        // 因为channel是带缓冲的，所以这里有可能会阻塞
        // 允许同时存在的goroutine的数量就是channel的缓存长度
        s.c <- struct{}{}
    }
}

func (s *Semaphore) Done() {
    <-s.c
    s.wg.Done()
}

func (s *Semaphore) Wait() {
    s.wg.Wait()
}
