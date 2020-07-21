package main

import (
	"wheel/gomicro-demo/protos"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
)

/*对外提供http服务，本身是一个微服务，在http内部调用grpc微服务
 */

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 客户端也可以作为服务，爱注册不注册，看是否需要，这里就不注册了
	grpcService := micro.NewService(
		micro.Name("productservice.client"),
	)
	// 这一步其实就是连接到grpc server
	prodService := protos.NewProductService("productservice", grpcService.Client())

	// 这个web server是grpc服务的客户端，http对外提供服务，并不是非得进行服务注册，看需要
	httpServer := web.NewService(
		web.Name("http_prod_service"),
		web.Address(":8001"),
		web.Handler(NewGinRouter(prodService)),
		web.Registry(etcdReg),
	)

	httpServer.Init()
	httpServer.Run()
}
