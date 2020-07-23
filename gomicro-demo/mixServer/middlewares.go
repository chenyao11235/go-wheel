package main

import (
	"fmt"
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

//ErrorMiddleware 错误拦截中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(500, gin.H{
					"status": fmt.Sprintf("%s", r),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
