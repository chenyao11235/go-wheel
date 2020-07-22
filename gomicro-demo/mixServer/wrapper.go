package main

import (
	"context"
	"fmt"

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
