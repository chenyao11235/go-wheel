package cyrpc

import (
    "context"
    "os"
    "os/signal"
    "syscall"
)

type Server struct {
    opts     *ServerOptions     // 服务端的选项设置
    services map[string]Service // 都注册了哪些的服务
}

func (s *Server) Close() {

}

func (s *Server) Serve() {
    for _, service := range s.services {
        go service.Serve(s.opts)
    }

    // 监听系统信号，就是kill命令发出的信号
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSEGV)
    <-ch

    s.Close()
}

// 定义服务的接口实现
type Service interface {
    Register(string, Handler)
    Serve(*ServerOptions)
    Close()
}

type service struct {
    svr         interface{}
    ctx         context.Context
    cancel      context.CancelFunc
    serviceName string
    handlers    map[string]Handler
    opts        *ServerOptions
}

type Handler interface {
}

// 用于设置服务端的函数

func NewServer(opt ...ServerOption) *Server {
    s := &Server{
        opts:     &ServerOptions{},
        services: make(map[string]Service),
    }

    for _, o := range opt {
        o(s.opts)
    }

    return s
}
