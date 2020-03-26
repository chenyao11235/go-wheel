package cron

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int // 规定每次调度的任务数量
	longLived  bool
	Dispatcher fn // 调度器，函数类型
	Executor   fn // 执行器，函数类型
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		dataSize:   size,
		longLived:  longlived,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select {
		case c := <-r.Controller:
			// 调度器分发具体的任务数据
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}
			// 执行器获取具体的任务数据并执行具体操作
			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}

			}
			//发生异常之后，直接退出
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

func (r *Runner) start() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
