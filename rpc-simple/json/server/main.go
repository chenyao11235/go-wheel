package main

import (
    "errors"
    "fmt"
    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    "time"
)

// 参数
type Args struct {
    A int `useJsonFile: "a"`
    B int `useJsonFile: "b"`
}

// 商
type Quotient struct {
    Quo int `useJsonFile: "quo"`
    Rem int `useJsonFile: "rem"`
}

type Arith int

func (t *Arith) Multiply(a *Args, reply *int) error {
    *reply = a.A * a.B
    return nil
}

func (t *Arith) Divide(a *Args, quo *Quotient) error {
    if a.B == 0 {
        err := errors.New("divide by zero")
        return err
    }

    quo.Quo = a.A / a.B
    quo.Rem = a.A % a.B
    return nil
}

func main() {
    arith := new(Arith)
    rpc.Register(arith)

    listener, err := net.Listen("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatal("fatal error", err)
    }

    for {
        fmt.Printf("accept at %v\n", time.Now())
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        fmt.Printf("get at %v\n", time.Now())
        jsonrpc.ServeConn(conn)
    }
}
