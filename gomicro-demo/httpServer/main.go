package main

import (
	"fmt"
	"wheel/gomicro-demo/models"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
)

/*基于go-micro实现一个http的微服务
 */

//ProductsRequests 商品请求
type ProductsRequest struct {
	Size int `form:"size"`
}

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	ginRouter := gin.Default()

	v1 := ginRouter.Group("/v1")
	{
		v1.Handle("POST", "/prods", func(c *gin.Context) {
			var req ProductsRequest
			err := c.Bind(&req)
			if err != nil || req.Size < 0 {
				req = ProductsRequest{Size: 3}
			}
			fmt.Println(req)
			c.JSON(200, gin.H{"data": models.NewProductList(req.Size)})
		})
	}
	httpServer := web.NewService(
		web.Name("productservice"),
		// web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(etcdReg),
	)
	// 使用Init方法可以使得程序具备命令行的功能
	// 通过命令行的方式指定addr，而不是在NewService中
	// 命令行参数是 --server_address :8001
	httpServer.Init()
	httpServer.Run()
}
