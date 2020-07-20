package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"wheel/gomicro-demo/models"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	mhttp "github.com/micro/go-plugins/client/http"
)

// 测试 服务发现的 是否能够正常进行
func TestProductServiceDiscovery(t *testing.T) {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	services, err := etcdReg.GetService("productservice")
	if err != nil {
		t.Error(err)
	}
	// 随机获取的一个server address
	// 还支持其他的 负载均衡算法
	next := selector.Random(services)
	// 轮讯
	// next = selector.RoundRobin(services)
	node, err := next()
	if err != nil {
		t.Error(err)
	}
	t.Log(node.Id)
	t.Log(node.Address)
	t.Log(node.Metadata)

	serverAddr := node.Address
	req, _ := http.NewRequest("GET", "http://"+serverAddr+"/v1/prods", nil)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(content))
}

//使用go-micro-plugin调用服务
func TestGoMicroPlugin(t *testing.T) {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 自定义选择器
	mySelector := selector.NewSelector(
		selector.Registry(etcdReg),
		// 设置负载均衡的方式
		selector.SetStrategy(selector.RoundRobin),
	)

	myClient := mhttp.NewClient(
		client.Selector(mySelector),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("productservice", "/v1/prods",
		models.ProdsRequest{Size: 10})
	// var rsp map[string]interface{}
	// 使用在proto中定义的response
	var rsp models.ProdListResponse
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		t.Error(err)
	}
	t.Log(rsp.GetData())
}
