package ps

import (
    "sync"
    "time"
)

// Package pubsub implements a simple multi-topic pub-sub library.
type (
    subscriber chan interface{}
    topicFunc  func(v interface{}) bool // 主题为一个过滤器
)

// 发布者对象
type Publisher struct {
    m           sync.RWMutex // 读写锁
    buffer      int          // 订阅队列的缓存大小
    timeout     time.Duration
    subscribers map[subscriber]topicFunc
}
