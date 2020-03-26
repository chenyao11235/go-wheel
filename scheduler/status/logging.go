package status

import "github.com/mongodb/mongo-go-driver/mongo"

// 负责将日志信息存储到数据库中

type Logging struct {
	client *mongo.Client
	logChan chan *JobLog
}

func InitLogging() {
	// 先建立数据库链接
}
