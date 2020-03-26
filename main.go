package main

import (
    "fmt"
)

type MyM struct {
    i int64
}

func (c MyM) Sing() {
    fmt.Println("i am sing...")
}

type MyN struct {
    i int64
}

func (c MyN) Sing() {
    fmt.Println("i am sing...")
}

func Test() {
}

func main() {
    i := 1
    n := 100
    for i < n {
        i = i * 3
    }
}
