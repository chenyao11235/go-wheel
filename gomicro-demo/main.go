package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	ginRouter := gin.Default()
	ginRouter.Handle("GET", "/", func(c *gin.Context) {
		data := make([]interface{}, 0)
		c.JSON(200, gin.H{
			"data": data,
		})
	})

	web := web.NewService(
		web.Address(":8000"),
		web.Handler(ginRouter),
	)

	web.Run()
}
