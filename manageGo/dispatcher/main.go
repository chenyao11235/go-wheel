package main

import (
    "fmt"
    "time"
)

type Job interface {
    Do()
}

type Upload struct {
    num int
}

var jobQueue chan Job

func (u *Upload) Do() {
    fmt.Println(u.num)
    time.Sleep(1 * time.Second)
}

//
type Worker struct {
    JobQueue    chan Job
    WorkerQueue chan chan Job
    quit        chan bool
}

func NewWorker(workerQueue chan chan Job) Worker {
    return Worker{
        JobQueue:    make(chan Job),
        WorkerQueue: workerQueue,
        quit:        make(chan bool),
    }
}

func (w *Worker) Start() {
    go func() {
        for {
            w.WorkerQueue <- w.JobQueue
            select {
            case job := <-w.JobQueue:
                job.Do()
            }
        }
    }()
}

func (w *Worker) Stop() {
    go func() {
        w.quit <- true
    }()
}

type Dispatcher struct {
    maxWorkers int
    WorkerPool chan chan Job
}

func NewDispatcher(maxWorker int) *Dispatcher {
    pool := make(chan chan Job, maxWorker)
    return &Dispatcher{
        maxWorkers: maxWorker,
        WorkerPool: pool,
    }
}

func (d *Dispatcher) Run() {
    for i := 0; i < d.maxWorkers; i++ {
        worker := NewWorker(d.WorkerPool)
        worker.Start()
    }

    go d.dispatch()
}

func (d *Dispatcher) dispatch() {
    for {
        select {
        case job := <-jobQueue:
            go func(job Job) {
                jobQueue := <-d.WorkerPool
                jobQueue <- job
            }(job)
        }
    }
}

func main() {
    jobQueue = make(chan Job)

    d := NewDispatcher(100 * 100)
    d.Run()

    n := 100 * 100 * 100
    for i := 0; i < n; n++ {
        job := &Upload{
            n,
        }
        jobQueue <- job
    }

}
