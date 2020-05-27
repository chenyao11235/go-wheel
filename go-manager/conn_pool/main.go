package main

import (
    "io"
    "log"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

const (
    maxGoroutines    = 25
    poolResourceSize = 2
)

var (
    idCounter int32
)

// dbConnection 模拟要共享的资源
type dbConnection struct {
    ID int32
}

// 要实现io.Closer 接口
func (dbConn *dbConnection) Close() error {
    log.Println("Close: Connection", dbConn.ID)
    return nil
}

// 制造连接
func createConnection() (io.Closer, error) {
    id := atomic.AddInt32(&idCounter, 1)
    log.Println("Create New Connection: ", id)

    return &dbConnection{ID: id}, nil
}

// 测试连接的资源池
func performQueries(query int, p *Pool) {
    conn, err := p.Acquire()
    if err != nil {
        log.Println(err)
        return
    }

    defer p.Release(conn)

    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func main() {
    var (
        wg sync.WaitGroup
    )

    wg.Add(maxGoroutines)

    p, err := NewPool(createConnection, poolResourceSize)
    if err != nil {
        log.Println(err)
        return
    }

    for query := 0; query < maxGoroutines; query++ {
        go func(q int) {
            performQueries(q, p)
            wg.Done()
        }(query)
    }

    wg.Wait()
    log.Println("Shutdown Program...")
    p.Close()
}
