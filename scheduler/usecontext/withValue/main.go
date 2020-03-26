package main

import (
    "context"
    "fmt"
)

func main() {
    parent := context.Background()
    ctx := context.WithValue(parent, "key", "test")
    value := ctx.Value("key").(string)
    fmt.Println(value)
}
