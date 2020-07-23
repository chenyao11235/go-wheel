package main

import (
	"context"
	"strconv"
	"time"
	"wheel/gomicro-demo/protos"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

/*基于go-micro实现一个grpc的微服务
 */

//ProductService 基于grpc实现的商品服务
type ProductService struct {
}

//GetProdList  获取商品列表
func (s *ProductService) GetProdList(ctx context.Context, req *protos.ProdsRequest, rsp *protos.ProdListResponse) error {
	time.Sleep(time.Second * 3)
	prods := make([]*protos.ProdModel, 0)
	var i int32
	for i = 0; i < req.Size; i++ {
		prods = append(prods, &protos.ProdModel{
			ID:   100 + i,
			Name: "pname" + strconv.Itoa(100+int(i)),
		})
	}
	rsp.Data = prods
	return nil
}

//GetProdDetail 获取商品详细
func (s *ProductService) GetProdDetail(ctx context.Context, req *protos.ProdDetailRequest, rsp *protos.ProdDetailResponse) error {
	time.Sleep(time.Second * 3)
	id := req.GetID()
	prod := &protos.ProdModel{
		ID:   id,
		Name: "pname" + strconv.Itoa(int(id)),
	}
	rsp.Data = prod
	return nil
}

func main() {
	registry := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	grpcServer := micro.NewService(
		micro.Name("productservice"),
		micro.Address(":8011"),
		micro.Registry(registry),
	)

	grpcServer.Init()
	protos.RegisterProductServiceHandler(grpcServer.Server(), new(ProductService))
	grpcServer.Run()
}
