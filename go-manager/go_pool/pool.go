package main

import (
    "errors"
    "log"
    "sync/atomic"
)

type Task struct {
    Handler func(...interface{})
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
        capacity: capacity,
        state:    RUNNING,
        taskC:    make(chan *Task, capacity),
        closeC:   make(chan bool),
    }, nil
}

// 正在工作的goroutine 计数器 +1
func (p *Pool) incRunning() {
    atomic.AddUint64(&p.runningWorker, 1)
}

// 正在工作的goroutine 计数器 -1
func (p *Pool) decRunning() {
    atomic.AddUint64(&p.runningWorker, ^uint64(0))
}

// 获取当前正在工作的goroutine 数量
func (p *Pool) GetRunningWorkers() uint64 {
    return atomic.LoadUint64(&p.runningWorker)
}

// 获取携程池的容量
func (p *Pool) GetCap() uint64 {
    return atomic.LoadUint64(&p.capacity)
}

func (p *Pool) run() {
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
                task.Handler()
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

    if p.GetRunningWorkers() < p.capacity {
        p.run()
    }
    p.taskC <- task
    return nil
}

func (p *Pool) Close() {
    p.state = STOPED
    // 如果任务队列中还有任务，要继续执行完才能退出
    for len(p.taskC) > 0 {
    }

    p.closeC <- true
    close(p.taskC)
}
