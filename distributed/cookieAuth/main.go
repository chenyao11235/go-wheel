package main

import "wheel/distributed"

// 采用一致性hash算法实现分布式验证身份验证, 检验用户的身份信息有没有被篡改
// ip，用户id 黑名单
// 单一用户访问速率限制

var hostArray = []string{"192.168.0.106", "192.168.0.106"}

func main() {
	consistentHash := distributed.NewConsistent()

	// 载入所有的节点信息
	for _, ip := range hostArray {
		consistentHash.Add(ip)
	}

}
