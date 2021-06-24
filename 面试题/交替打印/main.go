package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	//打印数字
	go func() {
		i := 0
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				letter <- true
			default:
				break
			}
		}
	}()

	//打印字母
	go func() {
		wg.Add(1)
		str := "ABCD"
		for {
			select {
			case <-letter:
				fmt.Print(string(str[0]))
				str = str[1:]
				if len(str) <= 0 {
					wg.Done()
					return
				}
				number <- true
			default:
				break
			}
		}
	}()

	number <- true
	wg.Wait()
}
