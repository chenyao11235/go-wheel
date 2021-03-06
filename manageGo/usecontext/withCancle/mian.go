package main

import (
    "context"
    "fmt"
    "math/rand"
    "time"
)

func downloadPicture(ctx context.Context, urls chan string) {
    go func() {
        select {
        case <-ctx.Done():
            fmt.Println("download pic stop")
            return
        case url := <-urls:
            fmt.Println("download page url = " + url)
        }
    }()
}

func getUrls(ctx context.Context, urls chan string) {
    go func() {
        select {
        case <-ctx.Done():
            fmt.Println("get urls stop")
            return
        default:
            num := rand.Int31()
            url := fmt.Sprintf("url=%d", num)
            urls <- url
            time.Sleep(time.Second)
        }
    }()
}

func main() {
    parent := context.Background()
    urls := make(chan string)
    ctx, cancel := context.WithCancel(parent)
    downloadPicture(ctx, urls)
    getUrls(ctx, urls)
    time.Sleep(time.Second * 5)
    cancel()
    time.Sleep(time.Second * 10)
}
