package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(5)
		}
		// 没有close会死锁
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
