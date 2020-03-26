package main

import (
	"github.com/pkg/errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// 实现一致性hash算法

// 如果真实的节点太少的话，会造成数据的分布不均，可以增加虚拟节点，提高数据的均衡性
// 每个虚拟节点其实是真实节点的副本，数据仍是存储在真实节点的

type units []uint32

func (x units) Len() int {
	return len(x)
}

func (x units) Less(i, j int) bool {
	return x[i] > x[j]
}

func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type Consistent struct {
	circle         map[uint32]string // 保存着各个节点的IP地址
	virtualNodeNum int               // 虚拟节点的个数
	sortedHashes   units             // 对节点进行排序是为了查找节点的时候更加快速
	sync.RWMutex                     // map不是线程安全的需要枷锁
}

func NewConsistent() *Consistent {
	return &Consistent{
		circle:         make(map[uint32]string),
		virtualNodeNum: 20, // 每个真实节点设置20个虚拟节点，相当于有20个副本
	}
}

// 添加节点
func (c *Consistent) Add(element string) {
	c.Lock()
	defer c.Unlock()

	// 为添加的真实节点设置虚拟节点（副本）
	for i := 0; i < c.virtualNodeNum; i++ {
		c.circle[c.hashKey(c.generateKey(element, i))] = element
	}

	// 更新排序
	c.updateSortedHashes()
}

// 删除节点
func (c *Consistent) Remove(element string) {
	for i := 0; i < c.virtualNodeNum; i++ {
		delete(c.circle, c.hashKey(c.generateKey(element, i)))
	}
	c.updateSortedHashes()
}

// 获取最近的节点信息
func (c *Consistent) GetNode(element string) (ip string, err error) {
	var (
		key   uint32
		index int
	)

	c.RLock()
	defer c.RUnlock()

	if len(c.circle) == 0 {
		ip = ""
		err = errors.New("empty circle")
		return
	}

	key = c.hashKey(element)
	index = c.search(key)

	ip = c.circle[c.sortedHashes[index]]

	return
}

// 生成虚拟节点的key值
func (c *Consistent) generateKey(element string, index int) string {
	return element + strconv.Itoa(index)
}

// 生成节点的hash值
func (c *Consistent) hashKey(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}

	return crc32.ChecksumIEEE([]byte(key))

}

// 更新排序
func (c *Consistent) updateSortedHashes() {
	var (
		hashes units
	)

	hashes = c.sortedHashes[:0]
	// 判断切片容量，是否过大，如果过大就重置
	if cap(c.sortedHashes)/(c.virtualNodeNum*4) > len(c.circle) {
		hashes = nil
	}

	for k := range c.circle {
		hashes = append(hashes, k)
	}

	sort.Sort(hashes)

	c.sortedHashes = hashes

}

// 顺时针查找最近的节点
func (c *Consistent) search(key uint32) (index int) {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}

	// 使用二分查找
	index = sort.Search(len(c.sortedHashes), f)

	if index >= len(c.sortedHashes) {
		index = 0
	}

	return
}
