## 优雅地调度goroutine
#### close
- 对于启动的goroutine要思考它会不会被阻塞，如果不需要它了就应该即使关闭它
---

#### limit 限制请求速率
- 使用goolang.org/x/time/rate包
- 令牌桶算法

#### runner 
- 可以注册多个任务，由单独的goroutine顺序执行这些任务
- main goroutine控制task goroutine的超时时间，超时中断 
---

#### pool
- 使用有缓冲的chan实现资源池
- 管理在不同goroutine之间共享及独立使用的资源
- 实现资源的申请和回收
---

#### worker
- 无缓冲chan实现的goroutine池
- 控制并发数量
##### 适用场景：
有一堆任务等着执行，想要通过并发加快执行速度，但是又想要控制并发的数量

#### control 控制goroutine并发数量的几种方法

---
#### pc
生产者消费者模型

---
#### ps
发布者订阅者模型

---
#### cron

---
#### status
状态模式，保存每个goroutine的运行状态，可以中断运行中的goroutine

#### dispatcher
实现百万级并发的线程池

#### usecontext 
使用的context的方式

#### recycle
实现一个可以回收goroutine的线程池，相同原理的还有国人开发的开源项目ants