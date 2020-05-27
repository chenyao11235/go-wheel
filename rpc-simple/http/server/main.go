package main

import (
    "errors"
    "fmt"
    "net/http"
    "net/rpc"
)

// 参数
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

func main() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()

    err := http.ListenAndServe(":1234", nil)

    if err != nil {
        fmt.Println(err)
    }
}
