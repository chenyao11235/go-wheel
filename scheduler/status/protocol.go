package status

import (
	"context"
	"time"
)

// 定义结构体，工具函数等

// 定时任务
type Job struct {
	Name    string `json:"name"`    //  任务名
	Command string `json:"command"` // shell命令
}

// 变化事件
type JobEvent struct {
	EventType int //  SAVE, DELETE
	Job       *Job
}

// 任务执行状态
type JobExecuteInfo struct {
	Job        *Job               // 任务信息
	PlanTime   time.Time          // 理论上的调度时间
	RealTime   time.Time          // 实际的调度时间
	CancelCtx  context.Context    // 任务command的context
	CancelFunc context.CancelFunc //  用于取消command执行的cancel函数
}

// 任务调度计划
type JobSchedulePlan struct {
	Job      *Job      // 要调度的任务信息
	NextTime time.Time // 下次调度时间
}

// 任务执行结果
type JobExecuteResult struct {
	ExecuteInfo *JobExecuteInfo // 执行状态
	Output      []byte          // 脚本输出
	Err         error           // 脚本错误原因
	StartTime   time.Time       // 启动时间
	EndTime     time.Time       // 结束时间
}

// 构造任务执行计划
func BuildJobSchedulePlan(job *Job) (jobSchedulePlan *JobSchedulePlan, err error) {
	// 此函数中可以根据具体情况写相应的操作

	// 生成任务调度计划对象
	jobSchedulePlan = &JobSchedulePlan{
		Job: job,
		// 其他信息
	}
	return
}

// 构造执行状态信息
func BuildJobExecuteInfo(jobSchedulePlan *JobSchedulePlan) (jobExecuteInfo *JobExecuteInfo) {
	jobExecuteInfo = &JobExecuteInfo{
		Job:      jobSchedulePlan.Job,
		PlanTime: jobSchedulePlan.NextTime, // 计算调度时间
		RealTime: time.Now(),               // 真实调度时间
	}
	jobExecuteInfo.CancelCtx, jobExecuteInfo.CancelFunc = context.WithCancel(context.TODO())
	return
}

// 任务执行日志
type JobLog struct {
	JobName      string `json:"jobName" bson:"jobName"`           // 任务名字
	Err          string `json:"err" bson:"err"`                   // 错误原因
	Output       string `json:"output" bson:"output"`             // 脚本输出
	PlanTime     int64  `json:"planTime" bson:"planTime"`         // 计划开始时间
	ScheduleTime int64  `json:"scheduleTime" bson:"scheduleTime"` // 实际调度时间
	StartTime    int64  `json:"startTime" bson:"startTime"`       // 任务执行开始时间
	EndTime      int64  `json:"endTime" bson:"endTime"`           // 任务执行结束时间
}
