package main

import (
	"wheel/gomicro-demo/protos"

	"github.com/gin-gonic/gin"
)

//InitMiddleware 设置中间件
func InitMiddleware(prodService protos.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodservice"] = prodService
		c.Next()
	}
}
