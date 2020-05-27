package main

import (
    "errors"
    "fmt"
    "net"
    "net/rpc"
    "os"
    "time"
)

type Args struct {
    A, B int
}

// 商
type Quotient struct {
    Quo, Rem int
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

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}

func main() {
    arith := new(Arith)
    rpc.Register(arith)

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        fmt.Printf("accept at %v\n", time.Now())
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        fmt.Printf("get at %v\n", time.Now())
        rpc.ServeConn(conn)
    }
}
