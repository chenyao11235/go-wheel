package cron

// 在这个架构中调度器的作用就是从数据库(或者mongo，redis等)中取出具体的数据，发送到chan中
func VideoClearDispatcher(dc dataChan) error {
	return nil
}

// 执行具体的动作
func VideoClearExecutor(dc dataChan) error {
	return nil
}
