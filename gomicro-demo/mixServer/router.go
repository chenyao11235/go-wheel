package main

import (
	"wheel/gomicro-demo/protos"

	"github.com/gin-gonic/gin"
)

//NewGinRouter 新建gin
func NewGinRouter(prodService protos.ProductService) *gin.Engine {
	ginRouter := gin.Default()

	ginRouter.Use(InitMiddleware(prodService), ErrorMiddleware())

	v1 := ginRouter.Group("/v1")
	{
		v1.Handle("GET", "/prod/:id", GetProdDetail)
		v1.Handle("POST", "/prods", GetProdList)
	}
	return ginRouter
}
