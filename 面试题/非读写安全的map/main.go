package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

// 只有写锁，但是map并发读也是不安全的，建议使用读写锁

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	u := &UserAges{
		ages: map[string]int{},
	}
	u.Add("eric", 18)
	fmt.Println(u.Get("eric"))
}
