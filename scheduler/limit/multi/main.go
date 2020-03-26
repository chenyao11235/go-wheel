package multi

import (
    "context"
    "golang.org/x/time/rate"
    "log"
    "sort"
    "sync"
    "time"
)

// rate包显示速率只能精确到每秒钟，在这里我们可以定义每分钟，每小时的速率
// 会对速率的严格成都程度进行排序，最严格的那个最先判断

type RateLimiter interface {
    Wait(ctx context.Context) error
    Limit() rate.Limit
}

type multiLimiter struct {
    limiters []RateLimiter
}

func (ml *multiLimiter) Wait(ctx context.Context) error {
    for _, l := range ml.limiters {
        if err := l.Wait(ctx); err != nil {
            return err
        }
    }
    return nil
}

func (ml *multiLimiter) Limit() rate.Limit {
    return ml.limiters[0].Limit()
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
    byLimit := func(i, j int) bool {
        return limiters[i].Limit() < limiters[j].Limit()
    }
    sort.Slice(limiters, byLimit)

    return &multiLimiter{limiters: limiters}
}

func Per(eventCount int, duration time.Duration) rate.Limit {
    return rate.Every(duration / time.Duration(eventCount))
}

type APIConnection struct {
    rateLimiter *multiLimiter
}

func Open() *APIConnection {
    secondLimiter := rate.NewLimiter(Per(2, time.Second), 1)
    minuteLimiter := rate.NewLimiter(Per(100, time.Minute), 10)
    return &APIConnection{
        rateLimiter: MultiLimiter(secondLimiter, minuteLimiter),
    }
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
    if err := a.rateLimiter.Wait(ctx); err != nil {
        return err
    }
    return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
    if err := a.rateLimiter.Wait(ctx); err != nil {
        return err
    }
    return nil
}


func main() {
    apiConnection := Open()
    var wg = sync.WaitGroup{}
    wg.Add(20)

    for i := 0; i < 20; i++ {
        go func() {
            defer wg.Done()
            err := apiConnection.ReadFile(context.Background())
            if err != nil {
                log.Printf("Cannot readfile: %v", err)
            }
            log.Printf("ReadFile")
        }()
    }

    for i := 0; i < 20; i++ {
        go func() {
            defer wg.Done()
            err := apiConnection.ResolveAddress(context.Background())
            if err != nil {
                log.Printf("Cannot ResolveAddress: %v", err)
            }
            log.Printf("ResolveAddress")
        }()
    }
    wg.Wait()
}
