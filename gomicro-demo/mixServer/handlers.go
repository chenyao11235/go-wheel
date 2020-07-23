package main

import (
	"context"
	"strconv"
	"wheel/gomicro-demo/protos"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/gin-gonic/gin"
)

func newProd(id int32, pname string) *protos.ProdModel {
	return &protos.ProdModel{ID: id, Name: pname}
}

func defaultProds() (*protos.ProdListResponse, error) {
	prods := make([]*protos.ProdModel, 0)
	for i := 0; i < 5; i++ {
		prods = append(prods, newProd(10+int32(i), strconv.Itoa(10+i)))
	}

	res := &protos.ProdListResponse{}
	res.Data = prods
	return res, nil
}

//PanicError 直接抛出异常 不用仔判断err == nil
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

//GetProdDetail 获取单个商品详细
func GetProdDetail(c *gin.Context) {
	var req protos.ProdDetailRequest
	// 这里要断言一下
	prodService := c.Keys["prodservice"].(protos.ProductService)
	PanicError(c.BindUri(&req))
	// 这里就是调用的grpc的服务
	res, _ := prodService.GetProdDetail(context.Background(), &req)
	c.JSON(200, gin.H{
		"data": res.Data,
	})
}

//GetProdList 获取商品列表
func GetProdList(c *gin.Context) {
	var req protos.ProdsRequest
	// 这里要断言一下
	prodService := c.Keys["prodservice"].(protos.ProductService)
	PanicError(c.Bind(&req))
	// 这里就是调用的grpc的服务
	proRes, _ := prodService.GetProdList(context.Background(), &req)
	c.JSON(200, gin.H{
		"data": proRes.Data,
	})
}

//GetProdListWthHystrix 获取商品列表 使服务有熔断降级的功能
// hystrix其实有集成到go-micro的plugin, 这里先使用原生的hystrix-go来实现
func GetProdListWthHystrix(c *gin.Context) {
	var req protos.ProdsRequest
	// 这里要断言一下
	prodService := c.Keys["prodservice"].(protos.ProductService)
	err := c.Bind(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
	} else {
		config := hystrix.CommandConfig{
			Timeout: 1000,
		}

		hystrix.ConfigureCommand("getprods", config)
		var prodRes *protos.ProdListResponse
		err := hystrix.Do("getpros", func() error {
			// 这里就是调用的grpc的服务
			prodRes, err = prodService.GetProdList(context.Background(), &req)
			return err
		}, func(e error) error {
			prodRes, err = defaultProds()
			return err
		})

		if err != nil {
			c.JSON(500, gin.H{
				"status": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"data": prodRes.Data,
			})
		}
	}
}
