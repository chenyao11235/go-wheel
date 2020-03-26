package main

import (
    "fmt"
    "github.com/pkg/errors"
    "wheel/error/paras"
)

func ReturnError() paras.PaginationError {
    return paras.NewNegetivePageTokenError()
}

func main() {
    err := ReturnError()
    switch errors.Cause(err).(type) {
    case paras.PaginationError:
        fmt.Printf("参数%s出现错误 %s", err.Field(), err.Error())
    }
}
