package worker

import (
    "sync"
)

type Worker interface {
    Task()
}

type Pool struct {
    worker chan Worker
    sync.WaitGroup
}

// 提交工作到工作池
func (p *Pool) Run(w Worker) {
    p.worker <- w
}

// shutdown 等待所有goroutine停止工作
func (p *Pool) Shutdown() {
    close(p.worker)
    p.Wait()
}

// 创建一个新的工作池
func New(maxGoroutines int) *Pool {
    p := Pool{
        worker: make(chan Worker),
    }

    p.Add(maxGoroutines)

    for i := 0; i < maxGoroutines; i++ {
        go func() {
            for r := range p.worker {
                r.Task()
            }
            p.Done()
        }()
    }

    return &p
}
