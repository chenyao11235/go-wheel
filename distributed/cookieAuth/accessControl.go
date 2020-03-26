package main

import "sync"

//
type AccessControl struct {
	data map[int]interface{}
	sync.RWMutex
}

// 获取新记录
func (a *AccessControl) GetRecord(uid int) interface{} {
	a.Lock()
	defer a.Unlock()
	return a.data[uid]
}

// 设置新记录
func (a *AccessControl) SetNewRecord(uid int) {
	a.Lock()
	a.data[uid] = "hello world"
	defer a.Unlock()
}

func (a *AccessControl) GetDistributedRight() {

}
