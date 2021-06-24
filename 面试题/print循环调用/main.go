package main

import (
	"fmt"
	"sync/atomic"
)

type People struct {
	Name string
}

func (p *People) String() string {
	//People实现了fmt中的Stringer接口， 所以当把p作为fmt.Sprintf的拼接对象的时候会调用String()方法，所以产生了循环调用
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	fmt.Println(p)
}
