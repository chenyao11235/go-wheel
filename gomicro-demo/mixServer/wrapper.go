package main

import (
	"context"
	"fmt"
	"strconv"
	"wheel/gomicro-demo/protos"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
)

// 这个东西就是go-micro的类似于gin中的middleware机制

type logWrapper struct {
	client.Client
}

//Call 调用
func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

//NewLogWrapper 新建
func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

//ProdsWrapper GetProdList的熔断功能的封装
type ProdsWrapper struct {
	client.Client
}

func defaultWrapProds(rsp interface{}) {
	prods := make([]*protos.ProdModel, 0)
	for i := 0; i < 5; i++ {
		prods = append(prods, newProd(10+int32(i), strconv.Itoa(10+i)))
	}

	res := rsp.(*protos.ProdListResponse)
	res.Data = prods
}

//DefaultData 通用降级数据 适用用于GetProdList和GetProdDetail
func DefaultData(rsp interface{}) {
	switch t := rsp.(type) {
	case *protos.ProdDetailResponse:
		t.Data = newProd(10, "降级商品")
	case *protos.ProdListResponse:
		defaultWrapProds(rsp)
	default:
	}
}

//Call 调用
func (w *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                1000, // 请求超时 毫秒级
		RequestVolumeThreshold: 2,    // 发生多少次降级之后进行百分比的计算
		ErrorPercentThreshold:  50,   // 触发熔断开关的 降级请求的百分比
		SleepWindow:            5000, // 多少时间之后重新尝试请求 正常服务
	}

	hystrix.ConfigureCommand(cmdName, config)

	return hystrix.Do(cmdName, func() error {
		return w.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		DefaultData(rsp)
		return nil
	})
}

//NewProdsWrapper 新建
func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{Client: c}
}
