package limiter

import "log"

type ConnLimiter struct {
	concurrent int
	bucket     chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrent: cc,
		bucket:     make(chan int, cc),
	}
}

// 获取链接，不过不能获取链接说明已经超过阈值了
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrent {
		return false
	}

	cl.bucket <- 1
	return true
}

// 释放链接
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connction coming: %d", c)
}
