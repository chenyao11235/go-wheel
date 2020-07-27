package main

import (
	"context"
	"fmt"
	"wheel/gomicro/proto"

	"github.com/micro/go-micro/v2"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

//GreeterServieHandler 服务实现
type GreeterServieHandler struct {
}

//Hello 方法
func (s *GreeterServieHandler) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Message = " 你好, " + req.Name
	return nil
}

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(reg),
	)

	proto.RegisterGreeterHandler(service.Server(), new(GreeterServieHandler))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
