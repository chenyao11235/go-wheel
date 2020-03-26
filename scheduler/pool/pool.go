package pool

import (
    "errors"
    "io"
    "log"
    "sync"
)

var (
    ErrPoolClosed = errors.New("pool has been closed")
)

type Pool struct {
    resource chan io.Closer
    factory  func() (io.Closer, error)
    closed   bool
    sync.Mutex
}

// 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
    select {
    case r, ok := <-p.resource:
        log.Println("Acquired:", "Shared Resource")
        if !ok {
            return nil, ErrPoolClosed
        }
        return r, nil
    default:
        // 因为没有空闲资源，所以提供一个新资源
        log.Println("Acquired:", "New Resource")
        return p.factory()
    }
}

// 回收资源
func (p *Pool) Release(r io.Closer) {
    p.Lock()
    defer p.Unlock()

    if p.closed {
        r.Close()
        return
    }

    select {
    case p.resource <- r:
        log.Println("Release:", "In Queue")
    default:
        // 如果队列已满，则关闭这个资源
        log.Println("Release:", "Closing")
        r.Close()
    }

}

// 会让资源池停止工作，关闭所有资源
func (p *Pool) Close() {
    p.Lock()
    defer p.Unlock()

    if p.closed {
        return
    }

    p.closed = true
    close(p.resource)

    for r := range p.resource {
        r.Close()
    }
}

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
    if size < 0 {
        return nil, errors.New("size value too small")
    }
    return &Pool{
        factory:  fn,
        resource: make(chan io.Closer, size),
    }, nil
}
