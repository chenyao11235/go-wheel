package cron

import "time"

// 调度器： 将具体的任务数据发送到dataChan中
// 执行器： 从dataChan中获取具体的任务数据并执行

//特点： 该调度模式定时任务形式，调度器的作用就是取任务数据，当执行器把调度器取出来的任务执行完了之后，调度器再重新取任务，
//特点： 不会造成，一直取任务，造成任务堆了很多，却来不及执行
//特点： 和普通的定时轮询不一样的是当调度出来的任务全部执行完毕之后，可以立马再次调度，不会总是等待固定的时间

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func (w *Worker) start() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.start()
		}
	}
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval),
		runner: r,
	}
}

func main() {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.start()
}
