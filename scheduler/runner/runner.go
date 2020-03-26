package runner

import (
    "errors"
    "os"
    "os/signal"
    "time"
)

var (
    ErrTimeout   = errors.New("received timeout")
    ErrInterrupt = errors.New("received interrupt")
)

type Runner struct {
    interrupt chan os.Signal
    complete  chan error
    timeout   <-chan time.Time
    tasks     []func(int)
}

// 添加具体的任务函数
func (r *Runner) Add(tasks ...func(int)) {
    r.tasks = append(r.tasks, tasks...)
}

// 验证是否收到了中断信号
func (r *Runner) goInterrupt() bool {
    select {
    case <-r.interrupt:
        signal.Stop(r.interrupt)
        return true
    default:
        return false
    }
}

// 执行每一个任务函数
func (r *Runner) run() error {
    for id, task := range r.tasks {
        if r.goInterrupt() {
            return ErrInterrupt
        }
        task(id)
    }
    return nil
}

//  启动所有的任务，并监听通道
func (r *Runner) Start() error {
    // 接收系统所有的信号
    signal.Notify(r.interrupt, os.Interrupt)

    go func() {
        r.complete <- r.run()
    }()

    select {
    case err := <-r.complete:
        return err
    case <-r.timeout:
        return ErrTimeout
    }
}

func New(d time.Duration) *Runner {
    return &Runner{
        interrupt: make(chan os.Signal, 1),
        complete:  make(chan error),
        timeout:   time.After(d),
    }
}
