package status

import "time"

// 用于任务调度

type Scheduler struct {
	jobEventChan      chan *JobEvent              //  etcd任务事件队列
	jobPlanTable      map[string]*JobSchedulePlan // 任务调度计划表，保存在内存中
	jobExecutingTable map[string]*JobExecuteInfo  // 任务执行表，保存在内存中
	jobResultChan     chan *JobExecuteResult      // 任务结果队列
}

var (
	G_Scheduler *Scheduler
)

func (scheduler *Scheduler) handleJobEvent(jobEvent *JobEvent) {
	var (
		err             error
		jobSchedulePlan *JobSchedulePlan
		jobExecuteInfo  *JobExecuteInfo
		jobExecuting    bool
		jobExisted      bool
	)

	switch jobEvent.EventType {
	// 新增，修改任务信息
	case JOB_EVENT_SAVE:
		if jobSchedulePlan, err = BuildJobSchedulePlan(jobEvent.Job); err != nil {
			return
		}
		// 更新内存内存中任务表
		scheduler.jobPlanTable[jobEvent.Job.Name] = jobSchedulePlan
		//删除某个任务
	case JOB_EVENT_DELETE:
		if jobSchedulePlan, jobExisted = scheduler.jobPlanTable[jobEvent.Job.Name]; jobExisted {
			// 删除内存中任务表中的任务
			delete(scheduler.jobPlanTable, jobEvent.Job.Name)
		}
		// 强杀某个正在执行的任务
	case JOB_EVENT_KILL:
		// 通过任务执行信息中保存的上下文函数，取消正在执行的任务
		if jobExecuteInfo, jobExecuting = scheduler.jobExecutingTable[jobEvent.Job.Name]; jobExecuting {
			jobExecuteInfo.CancelFunc()
		}
	}

}

func (scheduler *Scheduler) handleJobResult() {

}

// 遍历保存在内存中的任务表，酌情执行
func (scheduler *Scheduler) TrySchedule() (scheduleAfter time.Duration) {
	var (
		jobPlan  *JobSchedulePlan
		now      time.Time
		nearTime *time.Time // 最近一次任务调度的时间
	)

	// 如果任务表为空则睡眠
	if len(scheduler.jobPlanTable) == 0 {
	}

	now = time.Now()

	for _, jobPlan = range scheduler.jobPlanTable {
		// 在这里可以对任务做一些判断，是否要尝试更新任务
		scheduler.TryStartExecuteJob(jobPlan)
		// 通过在遍历所有的任务的时候可以计算出下一次应该调度的时间
	}

	scheduleAfter = (*nearTime).Sub(now)
	return

}

// 遍历保存在内存中的任务执行状态的表
func (scheduler *Scheduler) TryStartExecuteJob(jobPlan *JobSchedulePlan) {
	var (
		jobExecuteInfo *JobExecuteInfo
		jobExecuting   bool
	)

	// 如果这个任务的状态是正在执行中(如果任务在表中就说明在执行)，就跳过
	if jobExecuteInfo, jobExecuting = scheduler.jobExecutingTable[jobPlan.Job.Name]; jobExecuting {
		return
	}

	//将任务的执行信息保存在内存中
	jobExecuteInfo = BuildJobExecuteInfo(jobPlan)
	scheduler.jobExecutingTable[jobPlan.Job.Name] = jobExecuteInfo

}

func (scheduler *Scheduler) ScheduleLoop() {

}

func InitScheduler() (err error) {
	return
}
