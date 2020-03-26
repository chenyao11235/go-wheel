package main

import (
    "errors"
    "fmt"
    "log"
    "sync/atomic"
    "time"
)

type Task struct {
    Handler func(v ...interface{})
    Params  []interface{}
}

type Pool struct {
    capacity      uint64
    runningWorker uint64
    state         int64
    taskC         chan *Task
    closeC        chan bool
    PanicHandler  func(interface{})
}

var ErrInvalidPoolCap = errors.New("invalid pool cap")

const (
    RUNNING = 1
    STOPED  = 0
)

func NewPool(capacity uint64) (*Pool, error) {
    if capacity < 0 {
        return nil, ErrInvalidPoolCap
    }
    return &Pool{
        capacity:      capacity,
        runningWorker: RUNNING,
        taskC:         make(chan *Task, capacity),
        closeC:        make(chan bool),
    }, nil
}

func (p *Pool) incRunning() {
    atomic.AddUint64(&p.runningWorker, 1)
}

func (p *Pool) decRunning() {
    atomic.AddUint64(&p.runningWorker, ^uint64(0))
}

func (p *Pool) getRunningWorkers() uint64 {
    return atomic.LoadUint64(&p.runningWorker)
}

func (p *Pool) getCap() uint64 {
    return atomic.LoadUint64(&p.capacity)
}

func (p *Pool) Run() {
    p.incRunning()

    go func() {
        defer func() {
            p.decRunning()
            if r := recover(); r != nil {
                if p.PanicHandler != nil {
                    p.PanicHandler(r)
                } else {
                    log.Printf("Wroker panic: %s/n", r)
                }
            }
        }()
        for {
            select {
            case task, ok := <-p.taskC:
                if !ok { // 如果channel被关闭，停止运行
                    return
                }
                task.Handler(task.Params...)
            case <-p.closeC:
                return
            }
        }
    }()
}

var ErrPoolAlreadyClosed = errors.New("pool already close")

func (p *Pool) Put(task *Task) error {
    if p.state == STOPED {
        return ErrPoolAlreadyClosed
    }

    if p.getRunningWorkers() < p.capacity {
        p.Run()
    }
    p.taskC <- task
    return nil
}

func (p *Pool) Close() {
    p.state = STOPED
    for len(p.taskC) > 0 {
    }

    p.closeC <- true
    close(p.taskC)
}

func main() {
    pool, err := NewPool(10)
    if err != nil {
        panic(err)
    }

    for i := 0; i < 20; i++ {
        _ = pool.Put(&Task{
            Handler: func(v ...interface{}) {
                fmt.Println(v)
            },
            Params: []interface{}{i},
        })
    }
    time.Sleep(10 * time.Second)
}
