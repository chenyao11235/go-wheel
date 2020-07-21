package main

import (
	"context"
	"fmt"
	"wheel/gomicro-demo/protos"

	"github.com/gin-gonic/gin"
)

//GetProdList 获取商品列表
func GetProdList(c *gin.Context) {
	var req protos.ProdsRequest
	// 这里要断言一下
	prodService := c.Keys["prodservice"].(protos.ProductService)
	err := c.Bind(&req)
	fmt.Println(req)
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
	} else {
		// 这里就是调用的grpc的服务
		proRes, _ := prodService.GetProdList(context.Background(), &req)
		c.JSON(200, gin.H{
			"data": proRes.Data,
		})
	}
}
